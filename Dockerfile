# Use an official Golang image to build the Go interpreter
FROM golang:1.21 AS builder
WORKDIR /app
COPY . .

# Build the main.exe for the Flask backend
RUN go build -o main.exe main.go

# Build the Windows executable (om.exe) for download
ENV GOOS=windows
ENV GOARCH=amd64
ENV CGO_ENABLED=0
RUN go build -o om.exe main.go

# Use a Python image for the final container
FROM python:3.11-slim
WORKDIR /app

# Copy both executables and frontend code
COPY --from=builder /app/main.exe ./main.exe
COPY --from=builder /app/om.exe ./om.exe
COPY frontend ./frontend
COPY omexamples ./omexamples

# Install Flask
RUN pip install --no-cache-dir -r frontend/requirements.txt

# Expose the Flask port
EXPOSE 5000

# Set the working directory to frontend
WORKDIR /app/frontend

# Run the Flask app
CMD ["python", "app.py"] 