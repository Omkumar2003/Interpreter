# Use an official Golang image to build the Go interpreter
FROM golang:1.21 AS builder
WORKDIR /app
COPY . .
RUN GOOS=windows GOARCH=amd64 go build -o main.exe main.go

# Use a Python image for the final container
FROM python:3.11-slim
WORKDIR /app

# Copy the built Go binary and frontend code
COPY --from=builder /app/main.exe ./main.exe
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