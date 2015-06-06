//line repl.y:2
package main

import __yyfmt__ "fmt"

//line repl.y:3
import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

//line repl.y:15
type replSymType struct {
	yys        int
	numInt     int64
	operator   string
	identifier string
}

const INTEGER = 57346
const ADDITION_OP = 57347
const SUBTRACTION_OP = 57348
const MULTIPLICATION_OP = 57349
const DIVISION_OP = 57350
const ASSIGNMENT_OP = 57351
const IDENTIFIER = 57352
const NEWLINE = 57353

var replToknames = []string{
	"INTEGER",
	"ADDITION_OP",
	"SUBTRACTION_OP",
	"MULTIPLICATION_OP",
	"DIVISION_OP",
	"ASSIGNMENT_OP",
	"IDENTIFIER",
	"NEWLINE",
}
var replStatenames = []string{}

const replEofCode = 1
const replErrCode = 2
const replMaxDepth = 200

//line repl.y:82

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

//line yacctab:1
var replExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
}

const replNprod = 11
const replPrivate = 57344

var replTokenNames []string
var replStates []string

const replLast = 25

var replAct = []int{

	3, 8, 9, 10, 11, 10, 11, 7, 12, 13,
	14, 15, 16, 17, 8, 9, 10, 11, 4, 1,
	6, 2, 0, 0, 5,
}
var replPact = []int{

	14, -1000, -1000, -4, -1000, -1, -1000, -1000, 14, 14,
	14, 14, 14, -2, -2, -1000, -1000, 9,
}
var replPgo = []int{

	0, 21, 0, 20, 19,
}
var replR1 = []int{

	0, 4, 1, 2, 2, 2, 2, 2, 2, 2,
	3,
}
var replR2 = []int{

	0, 1, 2, 1, 1, 1, 3, 3, 3, 3,
	3,
}
var replChk = []int{

	-1000, -4, -1, -2, 4, 10, -3, 11, 5, 6,
	7, 8, 9, -2, -2, -2, -2, -2,
}
var replDef = []int{

	0, -2, 1, 0, 3, 4, 5, 2, 0, 0,
	0, 0, 0, 6, 7, 8, 9, 10,
}
var replTok1 = []int{

	1,
}
var replTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
}
var replTok3 = []int{
	0,
}

//line yaccpar:1

/*	parser for yacc output	*/

var replDebug = 0

type replLexer interface {
	Lex(lval *replSymType) int
	Error(s string)
}

const replFlag = -1000

func replTokname(c int) string {
	// 4 is TOKSTART above
	if c >= 4 && c-4 < len(replToknames) {
		if replToknames[c-4] != "" {
			return replToknames[c-4]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func replStatname(s int) string {
	if s >= 0 && s < len(replStatenames) {
		if replStatenames[s] != "" {
			return replStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func repllex1(lex replLexer, lval *replSymType) int {
	c := 0
	char := lex.Lex(lval)
	if char <= 0 {
		c = replTok1[0]
		goto out
	}
	if char < len(replTok1) {
		c = replTok1[char]
		goto out
	}
	if char >= replPrivate {
		if char < replPrivate+len(replTok2) {
			c = replTok2[char-replPrivate]
			goto out
		}
	}
	for i := 0; i < len(replTok3); i += 2 {
		c = replTok3[i+0]
		if c == char {
			c = replTok3[i+1]
			goto out
		}
	}

out:
	if c == 0 {
		c = replTok2[1] /* unknown char */
	}
	if replDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", replTokname(c), uint(char))
	}
	return c
}

func replParse(repllex replLexer) int {
	var repln int
	var repllval replSymType
	var replVAL replSymType
	replS := make([]replSymType, replMaxDepth)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	replstate := 0
	replchar := -1
	replp := -1
	goto replstack

ret0:
	return 0

ret1:
	return 1

replstack:
	/* put a state and value onto the stack */
	if replDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", replTokname(replchar), replStatname(replstate))
	}

	replp++
	if replp >= len(replS) {
		nyys := make([]replSymType, len(replS)*2)
		copy(nyys, replS)
		replS = nyys
	}
	replS[replp] = replVAL
	replS[replp].yys = replstate

replnewstate:
	repln = replPact[replstate]
	if repln <= replFlag {
		goto repldefault /* simple state */
	}
	if replchar < 0 {
		replchar = repllex1(repllex, &repllval)
	}
	repln += replchar
	if repln < 0 || repln >= replLast {
		goto repldefault
	}
	repln = replAct[repln]
	if replChk[repln] == replchar { /* valid shift */
		replchar = -1
		replVAL = repllval
		replstate = repln
		if Errflag > 0 {
			Errflag--
		}
		goto replstack
	}

repldefault:
	/* default state action */
	repln = replDef[replstate]
	if repln == -2 {
		if replchar < 0 {
			replchar = repllex1(repllex, &repllval)
		}

		/* look through exception table */
		xi := 0
		for {
			if replExca[xi+0] == -1 && replExca[xi+1] == replstate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			repln = replExca[xi+0]
			if repln < 0 || repln == replchar {
				break
			}
		}
		repln = replExca[xi+1]
		if repln < 0 {
			goto ret0
		}
	}
	if repln == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			repllex.Error("syntax error")
			Nerrs++
			if replDebug >= 1 {
				__yyfmt__.Printf("%s", replStatname(replstate))
				__yyfmt__.Printf(" saw %s\n", replTokname(replchar))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for replp >= 0 {
				repln = replPact[replS[replp].yys] + replErrCode
				if repln >= 0 && repln < replLast {
					replstate = replAct[repln] /* simulate a shift of "error" */
					if replChk[replstate] == replErrCode {
						goto replstack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if replDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", replS[replp].yys)
				}
				replp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if replDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", replTokname(replchar))
			}
			if replchar == replEofCode {
				goto ret1
			}
			replchar = -1
			goto replnewstate /* try again in the same state */
		}
	}

	/* reduction by production repln */
	if replDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", repln, replStatname(replstate))
	}

	replnt := repln
	replpt := replp
	_ = replpt // guard against "declared and not used"

	replp -= replR2[repln]
	replVAL = replS[replp+1]

	/* consult goto table to find next state */
	repln = replR1[repln]
	replg := replPgo[repln]
	replj := replg + replS[replp].yys + 1

	if replj >= replLast {
		replstate = replAct[replg]
	} else {
		replstate = replAct[replj]
		if replChk[replstate] != -repln {
			replstate = replAct[replg]
		}
	}
	// dummy call; replaced with literal code
	switch replnt {

	case 1:
		//line repl.y:42
		{
			fmt.Println("=", replS[replpt-0].numInt)
		}
	case 2:
		//line repl.y:49
		{
		}
	case 3:
		replVAL.numInt = replS[replpt-0].numInt
	case 4:
		//line repl.y:55
		{
			replVAL.numInt = repllex.(*replLex).env[replS[replpt-0].identifier]
		}
	case 5:
		replVAL.numInt = replS[replpt-0].numInt
	case 6:
		//line repl.y:60
		{
			replVAL.numInt = replS[replpt-2].numInt + replS[replpt-0].numInt
		}
	case 7:
		//line repl.y:64
		{
			replVAL.numInt = replS[replpt-2].numInt - replS[replpt-0].numInt
		}
	case 8:
		//line repl.y:68
		{
			replVAL.numInt = replS[replpt-2].numInt * replS[replpt-0].numInt
		}
	case 9:
		//line repl.y:72
		{
			replVAL.numInt = replS[replpt-2].numInt / replS[replpt-0].numInt
		}
	case 10:
		//line repl.y:78
		{
			repllex.(*replLex).env[replS[replpt-2].identifier] = replS[replpt-0].numInt
			replVAL.numInt = replS[replpt-0].numInt
		}
	}
	goto replstack /* stack new state and value */
}
