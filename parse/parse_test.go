package parse

import (
	"os"
	"testing"
)

func TestParseBrokenShouldErr(t *testing.T) {
	result, err := Parse("{{")

	if err == nil {
		t.Errorf("Expected error")
	}

	if result != nil {
		t.Errorf("Expected no result")
	}
}

func TestParseTextOnly(t *testing.T) {
	result, err := Parse("I HAVE NO LIQUID")

	if err != nil {
		t.Errorf("Expected no error")
	}

	if result == nil {
		t.Errorf("Expected result")
	}

	if len(result.Nodes) != 1 {
		t.Errorf("Expecting a single text node")
	}
}

func TestParseTagsOnly(t *testing.T) {
	result, err := Parse("{{'omg'}}")

	if err != nil {
		t.Errorf("Expected no error")
	}

	if result == nil {
		t.Errorf("Expected result")
	}

	if len(result.Nodes) > 2 {
		t.Errorf("Expecting two nodes")
	}
}

func TestParseRealFile(t *testing.T) {
	contents := make([]byte, 50000)
	file, err := os.Open("../example/products.liquid")

	if err != nil {
		panic(err)
	}

	file.Read(contents)
	result, err := Parse(string(contents))

	if err != nil {
		t.Errorf("Expected no error")
	}

	if result == nil {
		t.Errorf("Expected result")
	}

}
