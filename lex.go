package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

type stateFn func(*replLex) stateFn

type token struct {
	typ int
	val interface{}
}

type replLex struct {
	src    string
	start  int
	pos    int
	width  int
	state  stateFn
	tokens chan token
	env    map[string]int64
}

// The parser calls this method to get each new token.  This
// implementation returns operators and NUM.
func (x *replLex) Lex(yylval *replSymType) int {
	for {
		select {
		case token := <-x.tokens:
			switch token.typ {
			case INTEGER:
				yylval.numInt = token.val.(int64)
			case ADDITION_OP, SUBTRACTION_OP, MULTIPLICATION_OP, DIVISION_OP, ASSIGNMENT_OP:
				yylval.operator = token.val.(string)
			case IDENTIFIER:
				yylval.identifier = token.val.(string)
			case NEWLINE:
			}
			return token.typ
		default:
			x.state = x.state(x)
			if x.state == nil {
				return eof
			}
		}
	}
}

// The parser calls this method on a parse error.
func (x *replLex) Error(s string) {
	log.Printf("parse error: %s", s)
}

func NewReplLexer(src string, env map[string]int64) *replLex {
	return &replLex{src: src,
		tokens: make(chan token, 2),
		state:  blankState,
		env:    env,
	}
}

func (l *replLex) next() (r rune) {
	if l.pos >= len(l.src) {
		l.width = 0 // because backing up from having read eof should read eof again
		return 0
	}

	r, l.width = utf8.DecodeRuneInString(l.src[l.pos:])
	l.pos += l.width

	return r
}

func (l *replLex) unnext() {
	l.pos -= l.width
}

func (l *replLex) ignore() {
	l.start = l.pos
}

func (l *replLex) accept(valid string) bool {
	if strings.IndexRune(valid, l.next()) >= 0 {
		return true
	}
	l.unnext()
	return false
}

func (l *replLex) acceptRun(valid string) {
	for strings.IndexRune(valid, l.next()) >= 0 {
	}
	l.unnext()
}

func (l *replLex) acceptRunFunc(f func(rune) bool) {
	for f(l.next()) {
	}
	l.unnext()
}

func blankState(l *replLex) stateFn {
	switch r := l.next(); {
	case r == 0:
		return nil
	case unicode.IsDigit(r):
		return lexNumber
	case r == '+' || r == '-' || r == '*' || r == '/', r == '=':
		return lexOperator
	case r == '\n':
		return lexNewLine
	case isWhitespace(r):
		l.skipWhitespace()
		return blankState
	case isAlphanumeric(r):
		return lexIdentifier
	}
	return nil
}

func lexNumber(l *replLex) stateFn {
	l.acceptRun("0123456789")

	n, err := strconv.ParseInt(l.src[l.start:l.pos], 10, 64)
	if err != nil {
		fmt.Printf("%#v -- %#v", l, l.src[l.start:l.pos])
		panic("not an int")
	}

	l.tokens <- token{INTEGER, n}

	l.start = l.pos

	return blankState
}

func lexOperator(l *replLex) stateFn {
	t := token{val: l.src[l.start:l.pos]}
	switch t.val {
	case "+":
		t.typ = ADDITION_OP
	case "-":
		t.typ = SUBTRACTION_OP
	case "*":
		t.typ = MULTIPLICATION_OP
	case "/":
		t.typ = DIVISION_OP
	case "=":
		t.typ = ASSIGNMENT_OP
	default:
		panic("Unknown op")
	}
	l.tokens <- t
	l.start = l.pos
	return blankState
}

func lexIdentifier(l *replLex) stateFn {
	l.acceptRunFunc(isAlphanumeric)

	l.tokens <- token{IDENTIFIER, l.src[l.start:l.pos]}
	l.start = l.pos
	return blankState
}

func lexNewLine(l *replLex) stateFn {
	l.tokens <- token{NEWLINE, l.src[l.start:l.pos]}
	l.start = l.pos
	return blankState
}

func (l *replLex) skipWhitespace() {
	var r rune
	for r = l.next(); isWhitespace(r); r = l.next() {
	}

	if r != 0 {
		l.unnext()
	}

	l.ignore()
}

func isWhitespace(r rune) bool {
	return r != '\n' && unicode.IsSpace(r)
}

func isAlphanumeric(r rune) bool {
	return r == '_' || unicode.In(r, unicode.Letter, unicode.Digit)
}
