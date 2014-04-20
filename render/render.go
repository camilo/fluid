package render

import (
	"github.com/camilo/fluid/parse"
	"strings"
)

type Filter func(interface{}) interface{}

var filters map[string]Filter

type renderer struct {
	bindings map[string]interface{}
	filters  map[string]Filter
}

func (r *renderer) render(tree *parse.ListNode) string {
	var output []string

	for _, node := range tree.Nodes {
		output = append(output, node.Render(r.bindings))
	}

	return strings.Join(output, "")
}

func Render(tree *parse.ListNode, bindings map[string]interface{}) string {
	r := renderer{bindings: bindings}
	return r.render(tree)
}
