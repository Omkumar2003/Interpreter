# Interpreter Web App

A web-based playground and compiler for a custom programming language with Hindi-inspired keywords (`definekar`, `agar`, `warna`).

## Features
- Online REPL (web interface) to write and run code
- Downloadable Windows compiler (`om.exe`) to run `.om` files locally
- Example programs and language documentation
- Tutorial for using the compiler

## Try Online
Run the web REPL and learn the language at the homepage.

## Example Code
```om
definekar x = 10;
agar (x > 5) {
    x
} warna {
    0
}
```

## How to Run Locally (with Docker)
1. Build the Docker image:
   ```sh
   docker build -t interpreter-app .
   ```
2. Run the container:
   ```sh
   docker run -p 5000:5000 interpreter-app
   ```
3. Open [http://localhost:5000](http://localhost:5000) in your browser.

## Download the Compiler
- Go to the [Tutorial page](http://localhost:5000/tutorial) in the web app
- Download `om.exe` and run `.om` files on your Windows PC:
  ```sh
  om.exe filename.om
  ```

## Manual Setup (for Developers)
- Requires Go and Python (Flask)
- Build the Go interpreter:
  ```sh
  go build -o main.exe main.go
  ```
- Install Flask:
  ```sh
  pip install Flask
  ```
- Run the Flask app:
  ```sh
  cd frontend
  python app.py
  ```

## GitHub
[https://github.com/yourusername/interpreter](https://github.com/yourusername/interpreter)
