package symbolic

import "github.com/your-username/axiom/internal/ast"

// Expression represents a symbolic variable or operation derived from an AST node.
type Expression struct {
	Kind  string // e.g. "var", "binop", "literal"
	Value string
	Left  *Expression
	Right *Expression
}

// BuildExpression converts an AST node into a symbolic expression.
func BuildExpression(node *ast.Node) *Expression {
	if node == nil {
		return nil
	}

	switch node.Type {
	case "identifier":
		return &Expression{Kind: "var", Value: node.Content}
	case "number_literal", "integer_literal":
		return &Expression{Kind: "literal", Value: node.Content}
	case "binary_expression":
		return buildBinaryExpression(node)
	default:
		return &Expression{Kind: "unknown", Value: node.Content}
	}
}

func buildBinaryExpression(node *ast.Node) *Expression {
	expr := &Expression{Kind: "binop"}
	for _, child := range node.Children {
		switch child.Type {
		case "identifier", "number_literal":
			e := BuildExpression(child)
			if expr.Left == nil {
				expr.Left = e
			} else {
				expr.Right = e
			}
		default:
			expr.Value = child.Content // operator
		}
	}
	return expr
}
