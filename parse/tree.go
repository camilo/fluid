package parse

import (
	"fmt"
	"strings"
)

type nodeType int

func (n nodeType) String() string {

	switch n {
	case templateNode:
		return "templateNode"
	case textNode:
		return "textNode"
	case liquidNode:
		return "liquidNode"
	case outputSourceNode:
		return "outputSourceNode"
	case litNode:
		return "litNode"
	case identifierNode:
		return "identifierNode"
	case fieldNode:
		return "fieldNode"
	case filterListNode:
		return "filterListNode"
	}

	return "Unknown"

}

const (
	templateNode = iota
	textNode
	liquidNode
	outputSourceNode
	litNode
	identifierNode
	fieldNode
	filterListNode
)

type Renderable interface {
	Render() string
}

type parseNode struct {
	typ   nodeType
	left  *parseNode
	right *parseNode
	value interface{}
}

func (t *parseNode) String() string {
	return fmt.Sprintf("%s:[%s, %s] \n --- \n %s \n --- \n", t.typ, t.left,
		t.right, t.value)
}

func (t *parseNode) Render(bindings map[string]interface{}) string {
	typ := t.typ
	var output string
	switch typ {
	case textNode, litNode:
		output = t.value.(string)
	case liquidNode:
		input := ""

		if t.left.value.(*parseNode).typ == identifierNode {
			name := t.left.value.(*parseNode).value.(identifier).identifier
			bound_value := bindings[name]
			if bound_value == nil {
			} else {

				switch bound_value.(type) {
				case string:
					input = bindings[name].(string)
				case []string:
					// IRL keep it here, do not stringify
					input = strings.Join(bindings[name].([]string), ",")
				}
			}
		} else {
			output = output + t.left.value.(*parseNode).Render(bindings)
		}

		// Run filters then append
		if t.right != nil {
		} else {
			output = input + output
		}
	default:

		if t.left != nil {
			output = output + t.left.Render(bindings)
		}

		if t.right != nil {
			output = output + t.right.Render(bindings)
		}
	}

	return output
}

func (p *parseNode) appendNode(node *parseNode) {
	fmt.Printf("%s ", node)
}

type ListNode struct {
	Nodes []*parseNode
}

func (l *ListNode) appendNode(p *parseNode) {
	if p != nil {
		l.Nodes = append(l.Nodes, p)
	}
}

type identifier struct {
	identifier string
	field      string
}
