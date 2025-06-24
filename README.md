# Interpreter Web App

A web-based playground and compiler for a custom programming language with Hindi-inspired keywords (`definekar`, `agar`, `warna`).

## Features

* **Online REPL** - Write and run code directly in your browser
* **Code Viewer** - View, run, and download example programs
* **Language Learning** - Tutorial and documentation for the custom language
* **Windows Executable** - Download `om.exe` to run `.om` files locally
* **GitHub Integration** - Direct link to the source code repository

## Try Online

Visit the web REPL and learn the language at the homepage.

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

1. **Build the Docker image:**
   ```bash
   docker build -t interpreter-app .
   ```

2. **Run the container:**
   ```bash
   docker run -p 5000:5000 interpreter-app
   ```

3. **Open your browser at:** http://localhost:5000

## Download the Windows Executable

The Windows executable (`om.exe`) is automatically generated during the Docker build process.

### Download from Web App
1. Go to the homepage in the web app
2. Click "Download om.exe" in the "Run Locally" section
3. Save the file and run: `om.exe filename.om`

### Usage
```bash
om.exe filename.om
```

**Example:**
```bash
om.exe factorial.om
```

## Manual Setup (for Developers)

### Prerequisites
* Go 1.21+
* Python 3.7+ with Flask

### Steps
1. **Build the Go interpreter:**
   ```bash
   go build -o main.exe main.go
   ```

2. **Build the Windows executable:**
   ```bash
   GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -o om.exe main.go
   ```

3. **Install Flask:**
   ```bash
   pip install Flask
   ```

4. **Run the Flask app:**
   ```bash
   cd frontend
   python app.py
   ```

## Language Keywords

The language uses Hindi-inspired keywords:
- `definekar` - Variable declaration (instead of `let`)
- `agar` - If statement (instead of `if`)
- `warna` - Else statement (instead of `else`)
- `fn` - Function definition
- `return` - Return statement

## Built-in Functions

- `len(array)` - Get array length
- `first(array)` - Get first element
- `last(array)` - Get last element
- `rest(array)` - Get all elements except first
- `push(array, value)` - Add element to array
- `puts(value)` - Print value
- `index(array, i)` - Get element at index

## GitHub

View the source code and contribute: [https://github.com/Omkumar2003/Interpreter](https://github.com/Omkumar2003/Interpreter)

## About

A custom programming language interpreter built in Go with a web-based frontend for easy experimentation and learning.

<h1>Introduction</h1>
<p></p>This includes building a lexer, parser, abstract syntax tree representation,macro and an evaluator.
The primary objective is to create a fully functional interpreter for "Om" language , exclusively developed
within the scope of this project.

For building this projrct i took refrence from this book
https://interpreterbook.com/
Writing An Interpreter In Go by Thorsten Ball
</p>

<h1>How to Run it?</h1>
prerequisite:install golang
<h3>By repl</h3>
<ol>
<li>open terminal</li>
<li>type 'go run main.go'</li>
</ol>

<h3>By files</h3>
<ol><li>make sure your file extension ends with .om</li>
<li>to run file code ,open terminal and run 'go run main.go fileLocation' replace fileLoaction with actual file you want to run </li></ol>



<H1>Syntax</H1>

<ol>
  <li>
    let a = 120 ; <br>
    let b = "hello world" ;<br>
    let c = [12 , 13 , 14] ;<br>
    let d = [21 , 41 , 58 , "hello world"] ;<br>
    let e = [1, 2 * 2, 10 - 5, 8 / 2];<br>
  </li>
  <li>
    let add = fn(a, b) { a + b }; <br><br>
    let add = fn(a, b) { a + b };<br>
    let sub = fn(a, b) { a - b };<br>
    let applyFunc = fn(a, b, func) { func(a, b) };<br>
  </li>
  <li>
    puts is like print<br>
    let a = 10 ; <br>
    puts(a); <br><br>
    puts(10 * 20 / 45 -50);
  </li>
  <li>
    first gives first element of the array.
    <br>
    first(array_name);<br>
    <br>
    last gives last element of the array.
    <br>
    last(array_name);<br>
    <br>
    rest gives all element of the array other than first.
    <br>
    rest(array_name);<br>
    rest(rest(rest(rest(rest(a)))))<br>
    <br>
    push appends element to the array.
    <br>
    push(array_name, element);<br>
    <br>
    index gives the  element of the array at current index.
    <br>
    index(array_name, index_number);<br>
    <br>

  </li>
  <li>
    	let reverse = macro(a, b) { quote(unquote(b) - unquote(a)); };<br>
			reverse(2 + 2, 10 - 5);<br>
  </li>
  <li>
    let people = [{"name": "Alice", "age": 24}, {"name": "Anna", "age": 28}];<br>
    people[0]["name"];<br>

  </li>

</ol>
