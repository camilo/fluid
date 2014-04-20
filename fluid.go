package main

import (
	"fmt"
	"github.com/camilo/fluid/parse"
	"github.com/camilo/fluid/render"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Liquid file not provided")
		os.Exit(1)
	}

	filename := os.Args[1]
	contents := make([]byte, 50000)
	file, err := os.Open(filename)

	if err != nil {
		panic(err)
	}
	_, _ = file.Read(contents)

	tree, err := parse.Parse(string(contents))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Eror: %s", err)
	} else {
		bindings := map[string]interface{}{
			"description": "OMG OMG OMG OMG",
			"products":    []string{"product1", "product2"}}
		fmt.Printf(render.Render(tree, bindings))
	}
}
