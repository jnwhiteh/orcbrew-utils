package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"github.com/cespare/goclj/parse"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: %s [inputfile]\n", os.Args[0])
		os.Exit(2)
	}

	filename := os.Args[1]
	contentsBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed when reading file %s: %s", filename, err)
		os.Exit(2)
	}

	contents := string(contentsBytes)

	contents = regexp.MustCompile(`#:orcpub.dnd.e5{`).ReplaceAllString(contents, "{")
	contents = regexp.MustCompile(`#:orcpub.dnd.e5.character{`).ReplaceAllString(contents, "{")
	contents = regexp.MustCompile(`#:orcpub.dnd.e5.character`).ReplaceAllString(contents, "#")
	contents = regexp.MustCompile(`:orcpub.dnd.e5.character/`).ReplaceAllString(contents, "")
	contents = regexp.MustCompile(`orcpub.dnd.e5/`).ReplaceAllString(contents, "")
	contents = regexp.MustCompile(`orcpub.dnd.e5.[a-z-]*/`).ReplaceAllString(contents, "")

	buf := bytes.NewBufferString(contents)

	tt, err := parse.Reader(buf, "input.clj", 0)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing %s as Clojure: %s", filename, err)
		os.Exit(2)
	}

	fmt.Println(treeToJSON(tt))
}

func treeToJSON(tree *parse.Tree) string {
	return nodesToJSON(tree.Roots, 0)
}

func nodesToJSON(nodes []parse.Node, depth int) string {
	var buf bytes.Buffer
	for _, node := range nodes {
		buf.WriteString(strings.Repeat("  ", depth))
		buf.WriteString(nodeToJSON(node))
		buf.WriteString("\n")
	}

	return buf.String()
}

func nodeToJSON(node parse.Node) string {
	switch v := node.(type) {

	case *parse.KeywordNode:
		return fmt.Sprintf(`"%s"`, v.Val[1:])
	case *parse.StringNode:
		return fmt.Sprintf(`"%s"`, v.Val)
	case *parse.NumberNode:
		return v.Val
	case *parse.SymbolNode:
		return fmt.Sprintf(`"%s"`, v.Val)
	case *parse.SetNode:
		var vals []string
		for _, node := range filterNonValueNodes(v.Children()) {
			vals = append(vals, nodeToJSON(node))
		}

		return fmt.Sprintf("[%s]", strings.Join(vals, ","))
	case *parse.ListNode:
		var vals []string
		for _, node := range filterNonValueNodes(v.Children()) {
			vals = append(vals, nodeToJSON(node))
		}

		return fmt.Sprintf("[%s]", strings.Join(vals, ","))
	case *parse.VectorNode:
		var vals []string
		for _, node := range filterNonValueNodes(v.Children()) {
			vals = append(vals, nodeToJSON(node))
		}

		return fmt.Sprintf("[%s]", strings.Join(vals, ","))
	case *parse.MapNode:
		var keys []parse.Node
		var vals []parse.Node

		children := v.Children()
		children = filterNonValueNodes(children)

		if (len(children) % 2) != 0 {
			panic(v.String())
		}
		for idx, node := range children {
			if idx == 0 || (idx%2) == 0 {
				keys = append(keys, node)
			} else {
				// The value might be empty, remove the key in that case
				_, isNilValue := node.(*parse.NilNode)
				if isNilValue {
					keys = keys[0 : len(keys)-1]
				} else {
					vals = append(vals, node)
				}
			}
		}

		var entries []string
		for idx, key := range keys {
			var val = vals[idx]
			var keyString = nodeToJSON(key)
			if keyString[0] != '"' {
				keyString = fmt.Sprintf(`"%s"`, keyString)
			}
			entries = append(entries, fmt.Sprintf("%s: %s", keyString, nodeToJSON(val)))
		}

		return fmt.Sprintf("{%s}", strings.Join(entries, ","))
	case *parse.NewlineNode:
		return ""
	default:
		return node.String()
	}
}

func filterNonValueNodes(nodes []parse.Node) []parse.Node {
	var result []parse.Node

	for _, node := range nodes {
		switch v := node.(type) {
		case *parse.NewlineNode:
			break
		case *parse.CommentNode:
			break
		default:
			result = append(result, v)
		}
	}

	return result
}
