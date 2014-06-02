package render

import (
	"fmt"
	"github.com/camilo/fluid/parse"
)

func Example_liquid_only_with_binding() {
	liquid := "{{ sentence }}"
	bindings := map[string]interface{}{"sentence": "lorem ipsmum"}
	renderAndPrintLiquid(liquid, bindings)
	// Output:
	// lorem ipsmum
}

func Example_liquid_only_with_literal() {
	liquid := "{{ 'literal' }}"
	bindings := map[string]interface{}{}
	renderAndPrintLiquid(liquid, bindings)
	// Output:
	// 'literal'
}

func Example_text_only() {
	liquid := "This is just text "
	bindings := map[string]interface{}{}
	renderAndPrintLiquid(liquid, bindings)
	// Output:
	// This is just text
}

func renderAndPrintLiquid(contents string, bindings map[string]interface{}) {
	tree, err := parse.Parse(contents)
	if err != nil {
		panic(err)
	}
	fmt.Println(Render(tree, bindings))
}
