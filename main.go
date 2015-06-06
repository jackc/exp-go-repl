//go:generate -command yacc go tool yacc
//go:generate yacc -o repl.go -p "repl" repl.y
package main
