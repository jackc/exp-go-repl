package lex

import (
	"reflect"
	"testing"
)

func TestLex(t *testing.T) {
	tests := []struct {
		description string
		src         string
		expected    []token
	}{
		{"Empty source", "", []token{}},
		{"Integer", "42", []token{{"integer", int64(42)}}},
		{"Addition", "1 + 2", []token{{"integer", int64(1)}, {"operator", "+"}, {"integer", int64(2)}}},
		{"Subtraction", "3 - 2", []token{{"integer", int64(3)}, {"operator", "-"}, {"integer", int64(2)}}},
		{"Multiplation", "3 * 2", []token{{"integer", int64(3)}, {"operator", "*"}, {"integer", int64(2)}}},
		{"Division", "8 / 2", []token{{"integer", int64(8)}, {"operator", "/"}, {"integer", int64(2)}}},
		{"Operator without spaces", "1+2", []token{{"integer", int64(1)}, {"operator", "+"}, {"integer", int64(2)}}},
		{"Identifier", "foo", []token{{"identifier", "foo"}}},
		{"Assignment", "foo = bar", []token{{"identifier", "foo"}, {"operator", "="}, {"identifier", "bar"}}},
		{"Newline", "foo\nbar", []token{{"identifier", "foo"}, {"newline", "\n"}, {"identifier", "bar"}}},
	}

	for _, tt := range tests {
		lex := lexer{src: tt.src}
		lex.run()
		if !reflect.DeepEqual(tt.expected, lex.tokens) {
			t.Errorf("%s. Expected tokens (%#v), got tokens (%#v)", tt.description, tt.expected, lex.tokens)
		}
	}
}
