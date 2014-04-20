
//line fluid.y:2
package parse
import __yyfmt__ "fmt"
//line fluid.y:2
		import ("errors"
)

var ParseTree ListNode

//line fluid.y:9
type yySymType struct{
	yys int
  typ int
  pos int
  val string
  node *parseNode
}

const itemError = 57346
const itemIdentifier = 57347
const itemString = 57348
const itemNumberLit = 57349
const itemDotDot = 57350
const itemComparison = 57351
const itemLeftDelimLogic = 57352
const itemRightDelimLogic = 57353
const itemLeftDelimOutput = 57354
const itemRightDelimOutput = 57355
const itemPipe = 57356
const itemDot = 57357
const itemColon = 57358
const itemComma = 57359
const itemOpenSquare = 57360
const itemCloseSquare = 57361
const itemOpenRound = 57362
const itemCloseRound = 57363
const itemText = 57364
const itemEOF = 1
const itemSpace = 57366
const itemField = 57367
const itemNumber = 57368
const itemComplex = 57369

var yyToknames = []string{
	"itemError",
	"itemIdentifier",
	"itemString",
	"itemNumberLit",
	"itemDotDot",
	"itemComparison",
	"itemLeftDelimLogic",
	"itemRightDelimLogic",
	"itemLeftDelimOutput",
	"itemRightDelimOutput",
	"itemPipe",
	"itemDot",
	"itemColon",
	"itemComma",
	"itemOpenSquare",
	"itemCloseSquare",
	"itemOpenRound",
	"itemCloseRound",
	"itemText",
	"itemEOF",
	"itemSpace",
	"itemField",
	"itemNumber",
	"itemComplex",
}
var yyStatenames = []string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//line fluid.y:124


func Parse(liquid string) (*ListNode, error) {
  if 0 == yyParse(lex("lexer", liquid)) {
    return &ParseTree, nil
  } else {
    return nil, errors.New("")
  }

}





//line yacctab:1
var yyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 19
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 31

var yyAct = []int{

	14, 11, 12, 20, 4, 4, 10, 7, 18, 4,
	22, 17, 21, 16, 6, 3, 15, 19, 2, 9,
	5, 13, 8, 1, 0, 0, 0, 0, 0, 24,
	23,
}
var yyPact = []int{

	-7, -8, -15, -3, -5, -1000, -1000, -1000, -1000, 3,
	-1, -1000, -1000, -1000, -4, -1000, 12, -22, -2, -6,
	-1000, 12, -5, -2, -1000,
}
var yyPgo = []int{

	0, 23, 18, 6, 19, 8,
}
var yyR1 = []int{

	0, 1, 1, 1, 1, 1, 1, 1, 2, 3,
	3, 3, 3, 3, 5, 5, 5, 4, 4,
}
var yyR2 = []int{

	0, 0, 1, 1, 2, 2, 2, 2, 3, 1,
	1, 1, 1, 3, 1, 3, 3, 1, 3,
}
var yyChk = []int{

	-1000, -1, -2, 22, 12, -2, 22, 22, -2, -4,
	-3, 6, 7, 26, 5, 13, 14, 15, -5, 5,
	25, 14, 16, -5, -3,
}
var yyDef = []int{

	1, -2, 2, 3, 0, 6, 7, 5, 4, 0,
	17, 9, 10, 11, 12, 8, 0, 0, 18, 14,
	13, 0, 0, 16, 15,
}
var yyTok1 = []int{

	1, 23,
}
var yyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 0, 24, 25, 26, 27,
}
var yyTok3 = []int{
	0,
}

//line yaccpar:1

/*	parser for yacc output	*/

var yyDebug = 0

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

const yyFlag = -1000

func yyTokname(c int) string {
	// 4 is TOKSTART above
	if c >= 4 && c-4 < len(yyToknames) {
		if yyToknames[c-4] != "" {
			return yyToknames[c-4]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yylex1(lex yyLexer, lval *yySymType) int {
	c := 0
	char := lex.Lex(lval)
	if char <= 0 {
		c = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		c = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			c = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		c = yyTok3[i+0]
		if c == char {
			c = yyTok3[i+1]
			goto out
		}
	}

out:
	if c == 0 {
		c = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(c), uint(char))
	}
	return c
}

func yyParse(yylex yyLexer) int {
	var yyn int
	var yylval yySymType
	var yyVAL yySymType
	yyS := make([]yySymType, yyMaxDepth)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yychar := -1
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yychar), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yychar < 0 {
		yychar = yylex1(yylex, &yylval)
	}
	yyn += yychar
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yychar { /* valid shift */
		yychar = -1
		yyVAL = yylval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yychar < 0 {
			yychar = yylex1(yylex, &yylval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yychar {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error("syntax error")
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yychar))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yychar))
			}
			if yychar == yyEofCode {
				goto ret1
			}
			yychar = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		//line fluid.y:50
		{
	          ParseTree.appendNode(&parseNode{typ: -1})
	        }
	case 2:
		//line fluid.y:54
		{
	          ParseTree.appendNode(&parseNode{typ: -1, left: yyS[yypt-0].node})
	        }
	case 3:
		//line fluid.y:58
		{
	          text := &parseNode{typ: textNode, value: yyS[yypt-0].val}
	          ParseTree.appendNode(&parseNode{typ: -1, left: text})
	        }
	case 4:
		//line fluid.y:63
		{
	          text := &parseNode{typ: textNode, value: yyS[yypt-1].val}
	          ParseTree.appendNode(&parseNode{-1, text, yyS[yypt-0].node, nil})
	        }
	case 5:
		//line fluid.y:68
		{
	          text := &parseNode{typ: textNode, value: yyS[yypt-0].val}
	          ParseTree.appendNode(&parseNode{-1, yyS[yypt-1].node, text, nil})
	        }
	case 6:
		//line fluid.y:73
		{
	          ParseTree.appendNode(yyS[yypt-0].node)
	        }
	case 7:
		//line fluid.y:77
		{
	          text := &parseNode{typ: textNode, value: yyS[yypt-0].val}
	          ParseTree.appendNode(text)
	        }
	case 8:
		//line fluid.y:84
		{
	       // outputData rule returns liquidNode
       yyVAL.node = yyS[yypt-1].node
	       }
	case 9:
		//line fluid.y:91
		{yyVAL.node = &parseNode{typ: litNode, value: yyS[yypt-0].val}}
	case 10:
		//line fluid.y:93
		{yyVAL.node = &parseNode{typ: litNode, value: yyS[yypt-0].val}}
	case 11:
		//line fluid.y:95
		{yyVAL.node = &parseNode{typ: litNode, value:yyS[yypt-0].val}}
	case 12:
		//line fluid.y:97
		{yyVAL.node = &parseNode{typ: identifierNode, value: identifier{yyS[yypt-0].val, ""}}}
	case 13:
		//line fluid.y:99
		{yyVAL.node = &parseNode{typ: identifierNode, value: identifier{yyS[yypt-2].val, yyS[yypt-0].val}}}
	case 14:
		//line fluid.y:103
		{ //fmt.Printf("OMG: %s\n", $1)
            yyVAL.node=&parseNode{typ: filterListNode, value: yyS[yypt-0].val}}
	case 15:
		//line fluid.y:106
		{ //fmt.Printf("OMG: %s : %s \n", $1, $3.value)
            yyVAL.node=&parseNode{typ: filterListNode, value: yyS[yypt-2].val}}
	case 16:
		//line fluid.y:109
		{yyVAL.node=&parseNode{typ: filterListNode, value: nil}}
	case 17:
		//line fluid.y:113
		{
	        source := &parseNode{typ: outputSourceNode, value: yyS[yypt-0].node}
	        yyVAL.node = &parseNode{typ: liquidNode, left: source}
	      }
	case 18:
		//line fluid.y:118
		{
	        source := &parseNode{typ: outputSourceNode, value: yyS[yypt-2].node}
	        filters := &parseNode{typ: filterListNode, value: nil}
	        yyVAL.node = &parseNode{typ: liquidNode, left: source, right: filters}
	      }
	}
	goto yystack /* stack new state and value */
}
