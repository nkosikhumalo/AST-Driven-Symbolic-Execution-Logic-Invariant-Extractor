package ast

import (
	sitter "github.com/smacker/go-tree-sitter"
)

// Node is a simplified, language-agnostic AST node.
type Node struct {
	Type     string
	Content  string
	Children []*Node
}

// Walk traverses a tree-sitter tree and returns a language-agnostic Node tree.
func Walk(tree *ParsedTree) *Node {
	return walkNode(tree.Tree.RootNode(), tree.Source)
}

func walkNode(n *sitter.Node, source []byte) *Node {
	node := &Node{
		Type:    n.Type(),
		Content: n.Content(source),
	}

	for i := 0; i < int(n.ChildCount()); i++ {
		child := walkNode(n.Child(i), source)
		node.Children = append(node.Children, child)
	}

	return node
}
