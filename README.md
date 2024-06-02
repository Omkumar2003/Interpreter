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
