package lex

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

type stateFn func(*lexer) stateFn

type token struct {
	typ string
	val interface{}
}

type lexer struct {
	src    string
	start  int
	pos    int
	width  int
	tokens []token
}

func (l *lexer) run() {
	l.tokens = make([]token, 0)
	for state := blankState; state != nil; {
		state = state(l)
	}
}

func (l *lexer) next() (r rune) {
	if l.pos >= len(l.src) {
		l.width = 0 // because backing up from having read eof should read eof again
		return 0
	}

	r, l.width = utf8.DecodeRuneInString(l.src[l.pos:])
	l.pos += l.width

	return r
}

func (l *lexer) unnext() {
	l.pos -= l.width
}

func (l *lexer) ignore() {
	l.start = l.pos
}

func (l *lexer) accept(valid string) bool {
	if strings.IndexRune(valid, l.next()) >= 0 {
		return true
	}
	l.unnext()
	return false
}

func (l *lexer) acceptRun(valid string) {
	for strings.IndexRune(valid, l.next()) >= 0 {
	}
	l.unnext()
}

func (l *lexer) acceptRunFunc(f func(rune) bool) {
	for f(l.next()) {
	}
	l.unnext()
}

func blankState(l *lexer) stateFn {
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

func lexNumber(l *lexer) stateFn {
	l.acceptRun("0123456789")

	n, err := strconv.ParseInt(l.src[l.start:l.pos], 10, 64)
	if err != nil {
		fmt.Printf("%#v -- %#v", l, l.src[l.start:l.pos])
		panic("not an int")
	}

	l.tokens = append(l.tokens, token{"integer", n})

	l.start = l.pos

	return blankState
}

func lexOperator(l *lexer) stateFn {
	l.tokens = append(l.tokens, token{"operator", l.src[l.start:l.pos]})
	l.start = l.pos
	return blankState
}

func lexIdentifier(l *lexer) stateFn {
	l.acceptRunFunc(isAlphanumeric)

	l.tokens = append(l.tokens, token{"identifier", l.src[l.start:l.pos]})
	l.start = l.pos
	return blankState
}

func lexNewLine(l *lexer) stateFn {
	l.tokens = append(l.tokens, token{"newline", l.src[l.start:l.pos]})
	l.start = l.pos
	return blankState
}

func (l *lexer) skipWhitespace() {
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
