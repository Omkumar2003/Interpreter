package lexer

import (
	"github.com/om/interpreter/token"
)

type Lexer struct {
	input        string
	position     int  
	readPosition int  
	ch           byte 
}
//ye New function apni lifetime mein ek hi baar chal rha hai
//and it just mutates the lexer or more specifically lexer's input field
// ismein return  ek pointer ho rha hai , return krne ki koi jarrat nhi hti , kyuinki function sirf mutation ya sideffect kar rha hai
//but ye iska  func ka jo result hai wo hum repl.go mein save kr rhe hein  ,so that we can call nexttoken on it
func New(input string) *Lexer {
	 // it takes input as a parameter and then it makes an instance of lexer function in the memory then assign input key as input(parameter which which was given)
//read the below text about the pointer first
//in that i mentioned "In pointers if & is used before the variable name Then it gives the address of a variable name"
//but in the below case lexer is not variable name but it is a type declaraltion or struct // so how it is working
//struct literal
//In Go, you can create a new instance of a struct by using a struct literal. 
	l := &Lexer{input: input}
	l.readChar()
	//u can imagine it as engine start of the car...basically sets intial values
	return l
}
//In case of when you want to mutate variable value very frequently 
// then you should use pointer(pass by refrence) why?
//  if you don't use pointer then the only option left will be passed by value 
// the internal structure of possed by value is like this
//  it creates a temporary variable then Copy it Then mutate it then again copy it
//  this means that pass by value is costly and not memory efficient

// in go every go routine is isolated
//in a go routine there is a stack , inside that there is collection of frames (stack calls)
//every frame is isolated to each other
//this means variable defined in parent function should not be availabe as parameter or not be used in child function
//but we are are able to access the value // so what is happening
//in these cases it is passed as value .......this means less memory efficiency .more memory usage , more time
//to solve the problem we can use pointer// pointer can be is used to break isolation in a controlled way

// In pointers if & is used before the variable name Then it gives the address of a variable 
// in pointers * operator can be used in two ways 
// in first way if the star operator follows a Pointer name(pointer names stores address with the help of & operator) Then it will give the base value of the pointer
//  In second way if the star operator follows a type then it shows the base type would be that type

// (l *Lexer) if u only write (l lexer) then it will not change the value of the instance which u created in New function it just copy and make a new one pass by value
func (l *Lexer) readChar() {
	//The "readChar" method is used to get the next character from the input string and update the position and readPosition
 //loop nhi hai condition hai
 //if u call read char one time then it will read ony one time performs sideeffects
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhiteSpace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.EQ, Literal: literal}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.NOT_EQ, Literal: literal}
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case '[':
		tok = newToken(token.LBRACKET, l.ch)
	case ']':
		tok = newToken(token.RBRACKET, l.ch)
	case ':':
		tok = newToken(token.COLON, l.ch)
	case '"':
		tok.Type = token.STRING
		tok.Literal = l.readString()
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) skipWhiteSpace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readString() string {
	position := l.position + 1
	for {
		l.readChar()
		if l.ch == '"' || l.ch == 0 {
			break
		}
	}

	return l.input[position:l.position]
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
