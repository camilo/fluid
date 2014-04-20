package parse

import (
	"testing"
)

func TestEof(t *testing.T) {
	assertLexResult("EOF", "", []item{{itemEOF, 0, ""}}, t)
}

func TestTextOnly(t *testing.T) {
	input := "OMG ponies"
	expected := []item{
		{itemText, 0, "OMG ponies"},
		{itemEOF, 10, ""}}

	assertLexResult("text only", input, expected, t)
}

func TestTextWithOutputActionNumber(t *testing.T) {
	input := "OMG ponies {{things.morethings | things: 112.5 | lastfilter }} lol"
	expected := []item{
		{itemText, 0, "OMG ponies "},
		{itemLeftDelimOutput, 11, "{{"},
		{itemIdentifier, 13, "things"},
		{itemDot, 19, "."},
		{itemField, 20, "morethings"},
		{itemPipe, 31, "|"},
		{itemIdentifier, 33, "things"},
		{itemColon, 39, ":"},
		{itemNumber, 41, "112.5"},
		{itemPipe, 47, "|"},
		{itemIdentifier, 49, "lastfilter"},
		{itemRightDelimOutput, 60, "}}"},
		{itemText, 62, " lol"},
		{itemEOF, 66, ""}}

	assertLexResult("text and echo action", input, expected, t)
}

func TestTextWithOutputAction(t *testing.T) {
	input := "OMG ponies {{things.morethings | things: foo | lastfilter }} lol"
	expected := []item{
		{itemText, 0, "OMG ponies "},
		{itemLeftDelimOutput, 11, "{{"},
		{itemIdentifier, 13, "things"},
		{itemDot, 19, "."},
		{itemField, 20, "morethings"},
		{itemPipe, 31, "|"},
		{itemIdentifier, 33, "things"},
		{itemColon, 39, ":"},
		{itemIdentifier, 41, "foo"},
		{itemPipe, 45, "|"},
		{itemIdentifier, 47, "lastfilter"},
		{itemRightDelimOutput, 58, "}}"},
		{itemText, 60, " lol"},
		{itemEOF, 64, ""}}

	assertLexResult("text and echo action", input, expected, t)
}

func TestTextWithOutputActionAndStrings(t *testing.T) {
	input := "OMG ponies {{\"things.morethin\" | things: fo, 2 | lastfilter }} lol"
	expected := []item{
		{itemText, 0, "OMG ponies "},
		{itemLeftDelimOutput, 11, "{{"},
		{itemString, 13, "\"things.morethin\""},
		{itemPipe, 31, "|"},
		{itemIdentifier, 33, "things"},
		{itemColon, 39, ":"},
		{itemIdentifier, 41, "fo"},
		{itemComma, 43, ","},
		{itemNumber, 45, "2"},
		{itemPipe, 47, "|"},
		{itemIdentifier, 49, "lastfilter"},
		{itemRightDelimOutput, 60, "}}"},
		{itemText, 62, " lol"},
		{itemEOF, 66, ""}}

	assertLexResult("text and echo action", input, expected, t)
}
func TestTextWithOutputActionAndMissmatchedClosing(t *testing.T) {
	input := "OMG ponies {{things.morethings | things%} lol"
	expected := []item{
		{itemText, 0, "OMG ponies "},
		{itemLeftDelimOutput, 11, "{{"},
		{itemIdentifier, 13, "things"},
		{itemDot, 19, "."},
		{itemField, 20, "morethings"},
		{itemPipe, 31, "|"},
		{itemIdentifier, 33, "things"},
		{itemError, 39, "Unexpected closing logic tag: '%}' inside an output tag"},
	}

	assertLexResult("text and echo action", input, expected, t)
}

func TestTextWithLogicAction(t *testing.T) {
	input := "OMG ponies {%things   things%}"
	expected := []item{
		{itemText, 0, "OMG ponies "},
		{itemLeftDelimLogic, 11, "{%"},
		{itemRightDelimLogic, 28, "%}"},
		{itemEOF, 30, ""}}

	assertLexResult("text and no echo action", input, expected, t)
}

func TestTextWithOutputUnclosedAction(t *testing.T) {
	input := "OMG ponies {{things "
	expected := []item{
		{itemText, 0, "OMG ponies "},
		{itemLeftDelimOutput, 11, "{{"},
		{itemIdentifier, 13, "things"},
		{itemError, 20, "unclosed action"}}

	assertLexResult("unclosed echo action", input, expected, t)
}

func TestUnexpectedToken(t *testing.T) {
	input := "OMG ponies {{things ! "
	expected := []item{
		{itemText, 0, "OMG ponies "},
		{itemLeftDelimOutput, 11, "{{"},
		{itemIdentifier, 13, "things"},
		{itemError, 20, "unexpected token !"}}

	assertLexResult("unclosed echo action", input, expected, t)
}

func assertLexResult(lexerName string, input string, expected []item, t *testing.T) {
	var items []item
	lexer := lex(lexerName, input)

	for {
		item := lexer.nextItem()
		items = append(items, item)

		if item.typ == itemEOF || item.typ == itemError {
			break
		}

	}

	for i, it := range expected {
		if items[i].typ != it.typ || items[i].pos != it.pos || items[i].val !=
			it.val {
			t.Errorf("Expected type %s, pos %d, val '%s' got type %s pos %d val '%s'",
				it.typ, it.pos, it.val, items[i].typ, items[i].pos, items[i].val)
		}
	}

	if len(expected) != len(items) {
		t.Errorf("expected %d items got %d", len(expected), len(items))
	}
}
