%{

package main

import (
  "bufio"
  "fmt"
  "io"
  "log"
  "os"
)

%}

%union {
  numInt int64
  operator string
  identifier string
}

%type <numInt> line
%type <numInt> expr
%type <numInt> assignment

%token  <numInt> INTEGER
%token  <operator> ADDITION_OP
%token  <operator> SUBTRACTION_OP
%token  <operator> MULTIPLICATION_OP
%token  <operator> DIVISION_OP
%token  <operator> ASSIGNMENT_OP
%token  <identifier> IDENTIFIER
%token  <numInt> NEWLINE

%right ASSIGNMENT_OP
%left ADDITION_OP SUBTRACTION_OP
%left MULTIPLICATION_OP DIVISION_OP

%%

top:
  line
  {
    fmt.Println("=", $1)
  }

// Not sure why, but this thing fails unless there is a block for expr NEWLINE
line:
  expr NEWLINE
  {
  }

expr:
  INTEGER
| IDENTIFIER
  {
    $$ = repllex.(*replLex).env[$1]
  }
| assignment
| expr ADDITION_OP expr
  {
    $$ = $1 + $3
  }
| expr SUBTRACTION_OP expr
  {
    $$ = $1 - $3;
  }
| expr MULTIPLICATION_OP expr
  {
    $$ = $1 * $3;
  }
| expr DIVISION_OP expr
  {
    $$ = $1 / $3;
  }

assignment:
  IDENTIFIER ASSIGNMENT_OP expr
  {
    repllex.(*replLex).env[$1] = $3
    $$ = $3
  }
%%

// The parser expects the lexer to return 0 on EOF.  Give it a name
// for clarity.
const eof = 0

func main() {
  in := bufio.NewReader(os.Stdin)
  env := make(map[string]int64)
  for {
    if _, err := os.Stdout.WriteString("> "); err != nil {
      log.Fatalf("WriteString: %s", err)
    }
    line, err := in.ReadBytes('\n')
    if err == io.EOF {
      return
    }
    if err != nil {
      log.Fatalf("ReadBytes: %s", err)
    }

    lexer := NewReplLexer(string(line), env)

    replParse(lexer)
  }
}
