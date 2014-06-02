package parse

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"
)

type item struct {
	typ itemType
	pos Pos
	val string
}

type itemType int
type Pos int
type stateFn func(*lexer) stateFn

const itemTemplate = -2
const eof = -1

func (i item) String() string {
	return fmt.Sprintf("(%s %d '%s')", i.typ, i.pos, i.val)
}

func (i itemType) String() string {
	switch i {
	case itemEOF:
		return "EOF"
	case itemText:
		return "Text"
	case itemLeftDelimOutput:
		return "LeftDelimOuput"
	case itemRightDelimOutput:
		return "RightDelimOutput"
	case itemLeftDelimLogic:
		return "LeftDelimLogic"
	case itemRightDelimLogic:
		return "RightDelimLogic"
	case itemError:
		return "Error"
	case itemSpace:
		return "Space"
	case itemPipe:
		return "Pipe"
	case itemColon:
		return "Colon"
	case itemString:
		return "String"
	case itemField:
		return "Field"
	case itemNumber:
		return "Number"
	case itemIdentifier:
		return "Identifier"
	default:
		return "unknown"
	}
}

func (l *lexer) next() rune {
	if int(l.pos) >= len(l.input) {
		l.width = 0
		return eof
	}

	r, w := utf8.DecodeRuneInString(l.input[l.pos:])
	l.width = Pos(w)
	l.pos += l.width

	if r == '\n' {
		l.currentLine = l.currentLine + 1
	}
	return r
}

func (l *lexer) peek() rune {
	r := l.next()
	l.backup()
	return r
}

func (l *lexer) backup() {
	l.pos -= l.width
}

func (l *lexer) emit(t itemType) {
	l.items <- item{t, l.start, l.input[l.start:l.pos]}
	l.start = l.pos
}

func (l *lexer) ignore() {
	l.start = l.pos
}

func (l *lexer) accept(valid string) bool {

	if strings.IndexRune(valid, l.next()) >= 0 {
		return true
	}

	l.backup()
	return false

}

func (l *lexer) acceptRun(valid string) {
	for strings.IndexRune(valid, l.next()) >= 0 {
	}
	l.backup()
}

func (l *lexer) nextItem() item {
	item := <-l.items
	l.lastPos = item.pos
	return item
}

const (
	leftDelimLogic   = "{%"
	rightDelimLogic  = "%}"
	leftDelimOutput  = "{{"
	rightDelimOutput = "}}"
)

type lexer struct {
	name        string
	input       string
	state       stateFn
	pos         Pos
	start       Pos
	width       Pos
	lastPos     Pos
	items       chan item
	parenDepth  int
	nextDelim   itemType // this is suspicious might be better to keep the current action or no state at all if possible
	currentLine int
}

// Make go yacc happy
func (l *lexer) Lex(v *yySymType) (output int) {
	item := l.nextItem()

	if item.typ != itemEOF {
		v.val = item.val
		v.pos = int(item.pos)
		return int(item.typ)
	} else {
		return 0
	}
}

func (l *lexer) Error(e string) {
	fmt.Fprintf(os.Stderr, "Parser error: %s, current token: %s, current position %d, current line: %d", e, l.input[l.start:l.pos], l.pos, l.currentLine+1)
}

func (l *lexer) run() {
	for l.state = lexText; l.state != nil; {
		l.state = l.state(l)
	}
}

func (l *lexer) errorf(format string, args ...interface{}) stateFn {
	l.items <- item{itemError, l.start, fmt.Sprintf(format, args...)}
	return nil
}

func lex(name, input string) *lexer {
	l := &lexer{
		name:  name,
		input: input,
		items: make(chan item),
	}

	go l.run()
	return l
}

func LexAndPrint(input string) {
	lexer := lex("", input)
	for {
		item := lexer.nextItem()
		fmt.Printf("%v ", item)

		if item.typ == itemEOF || item.typ == itemError {
			break
		}

	}
}

func atAction(subStr string) bool {
	if strings.HasPrefix(subStr, leftDelimOutput) || strings.HasPrefix(subStr, leftDelimLogic) {
		return true
	}
	return false
}

func atLeftDelimLogic(subStr string) bool {
	if strings.HasPrefix(subStr, leftDelimLogic) {
		return true
	}
	return false
}

func atLeftDelimOutput(subStr string) bool {
	if strings.HasPrefix(subStr, leftDelimOutput) {
		return true
	}
	return false
}

func atRightDelimLogic(subStr string) bool {
	if strings.HasPrefix(subStr, rightDelimLogic) {
		return true
	}
	return false
}

func atRightlimOutput(subStr string) bool {
	if strings.HasPrefix(subStr, rightDelimOutput) {
		return true
	}
	return false
}

func atTerminator(r rune) bool {
	return r == ' ' || r == '\n' || r == '.' || r == ':' || r == '|' || r == '%' || r == '}' || r == ','
}

func isAlphaNumeric(r rune) bool {
	return r == '_' || unicode.IsLetter(r) || unicode.IsDigit(r)
}

// Can transition to either EOF or a delimiter
func lexText(l *lexer) stateFn {
	for {
		if atAction(l.input[l.pos:]) {
			if l.pos > l.start {
				l.emit(itemText)
			}

			return lexLeftDelim
		}

		if l.next() == eof {
			break
		}
	}

	if l.pos > l.start {
		l.emit(itemText)
	}
	l.emit(itemEOF)
	return nil
}

func lexLeftDelim(l *lexer) stateFn {
	if atLeftDelimLogic(l.input[l.pos:]) {
		l.pos += 2
		l.emit(itemLeftDelimLogic)
		l.nextDelim = itemRightDelimLogic
	} else if atLeftDelimOutput(l.input[l.pos:]) {
		l.pos += 2
		l.emit(itemLeftDelimOutput)
		l.nextDelim = itemRightDelimOutput
	} else {
		panic("unreachable")
	}

	l.parenDepth = 0
	return lexInsideAction
}

func lexInsideOutputAction(l *lexer) stateFn {
	for {

		if strings.HasPrefix(l.input[l.pos:], rightDelimOutput) {
			return lexRightDelimOutput
		}

		switch r := l.next(); {
		case r == '%':
			if l.peek() == '}' {
				return l.errorf("Unexpected closing logic tag: '%%}' inside an output tag")
			}
			fallthrough
		case r == eof:
			return l.errorf("unclosed action")
		case r == ' ':
			return lexSpace
		case r == '\'':
			return lexSingleQuoted
		case r == '"':
			return lexDoubleQuoted
		case r == ':': // Not sure if this can go inside an output {{, yes it can
			l.emit(itemColon)
		case r == '|':
			l.emit(itemPipe)
		case r == ',':
			l.emit(itemComma)
		case r == '.':
			if l.pos < Pos(len(l.input)) {
				r := l.input[l.pos]
				if r < '0' || '9' < r {
					return lexField
				}
			}
			fallthrough // '.' can start a number.
		case r == '+' || r == '-' || ('0' <= r && r <= '9'):
			l.backup()
			return lexNumber
		case isAlphaNumeric(r):
			l.backup()
			return lexIdentifier
		default:
			l.errorf("unexpected token %s", string(r))
		}
	}
}

func lexIdentifier(l *lexer) stateFn {
Loop:
	for {
		switch r := l.next(); {
		case isAlphaNumeric(r):
			// absorb.
		default:
			l.backup()
			word := l.input[l.start:l.pos]
			if !atTerminator(l.peek()) {
				return l.errorf("bad character %#U", r)
			}
			switch {
			case word[0] == '.':
				l.emit(itemField)
			default:
				l.emit(itemIdentifier)
			}
			break Loop
		}
	}
	return lexInsideAction
}

func lexField(l *lexer) stateFn {
	l.emit(itemDot)
	if atTerminator(l.peek()) {
		return lexInsideAction
	}
	var r rune
	for {
		r = l.next()
		if !isAlphaNumeric(r) {
			l.backup()
			break
		}
	}
	if !atTerminator(l.peek()) {
		return l.errorf("bad character %#U", r)
	}
	l.emit(itemField)
	return lexInsideAction
}

func lexNumber(l *lexer) stateFn {
	if !l.scanNumber() {
		return l.errorf("bad number syntax: %q", l.input[l.start:l.pos])
	}
	if sign := l.peek(); sign == '+' || sign == '-' {
		// Complex: 1+2i. No spaces, must end in 'i'.
		if !l.scanNumber() || l.input[l.pos-1] != 'i' {
			return l.errorf("bad number syntax: %q", l.input[l.start:l.pos])
		}
		l.emit(itemComplex)
	} else {
		l.emit(itemNumber)
	}
	return lexInsideAction
}

func (l *lexer) scanNumber() bool {
	// Optional leading sign.
	l.accept("+-")
	// Is it hex?
	digits := "0123456789"
	if l.accept("0") && l.accept("xX") {
		digits = "0123456789abcdefABCDEF"
	}
	l.acceptRun(digits)
	if l.accept(".") {
		l.acceptRun(digits)
	}
	if l.accept("eE") {
		l.accept("+-")
		l.acceptRun("0123456789")
	}
	// Is it imaginary?
	l.accept("i")
	// Next thing mustn't be alphanumeric.
	if isAlphaNumeric(l.peek()) {
		l.next()
		return false
	}
	return true
}

func lexQuote(l *lexer, limit byte) stateFn {
Loop:
	for {
		switch l.next() {
		case '\\':
			if r := l.next(); r != eof && r != '\n' {
				break
			}
			fallthrough
		case eof, '\n':
			return l.errorf("unterminated quoted string")
		case bytes.Runes([]byte{limit})[0]:
			break Loop
		}
	}
	l.emit(itemString)
	return lexInsideAction
}

func lexDoubleQuoted(l *lexer) stateFn {
	return lexQuote(l, '"')

}

func lexSingleQuoted(l *lexer) stateFn {
	return lexQuote(l, '\'')
}

func lexInsideLogicAction(l *lexer) stateFn {
	for {

		if strings.HasPrefix(l.input[l.pos:], rightDelimLogic) {
			//if l.parenDepth == 0 {
			return lexRightDelimLogic
			//}
			// error for parentesis
			// error for no echo action
		}

		switch r := l.next(); {
		case r == eof:
			return l.errorf("unclosed action")
		case r == ' ':
			return lexSpace
		case r == ':':
			l.emit(itemColon)
		case r == ',':
			l.emit(itemComma)
		case r == '|': // Not sure if this can go inside a logic {%
			l.emit(itemPipe)

		default:
			l.ignore() //change to error when the lexer is complete
		}
	}
}

func lexInsideAction(l *lexer) stateFn {

	if l.nextDelim == itemRightDelimOutput {
		return lexInsideOutputAction
	} else if l.nextDelim == itemRightDelimLogic {
		return lexInsideLogicAction
	} else {
		panic("unreachable")
	}

}

func lexSpace(l *lexer) stateFn {
	for ' ' == l.peek() {
		l.next()
	}
	l.ignore()
	return lexInsideAction(l)
}

func lexRightDelimOutput(l *lexer) stateFn {
	l.pos += Pos(len(rightDelimOutput))
	l.emit(itemRightDelimOutput)
	l.nextDelim = 0
	return lexText
}

func lexRightDelimLogic(l *lexer) stateFn {
	l.pos += Pos(len(rightDelimLogic))
	l.emit(itemRightDelimLogic)
	l.nextDelim = 0
	return lexText
}
