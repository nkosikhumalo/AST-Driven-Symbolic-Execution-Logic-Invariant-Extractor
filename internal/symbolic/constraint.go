package symbolic

import "github.com/your-username/axiom/internal/ast"

// PathConstraint represents a single conditional branch path.
type PathConstraint struct {
	Condition string // symbolic boolean expression
	Body      []string
}

// ExtractConstraints walks an AST node tree and maps conditional branches
// (if/else/switch) into a list of path constraints.
func ExtractConstraints(node *ast.Node) []PathConstraint {
	var constraints []PathConstraint
	collectConstraints(node, &constraints)
	return constraints
}

func collectConstraints(node *ast.Node, out *[]PathConstraint) {
	if node == nil {
		return
	}

	if node.Type == "if_statement" {
		pc := PathConstraint{
			Condition: extractCondition(node),
		}
		*out = append(*out, pc)
	}

	for _, child := range node.Children {
		collectConstraints(child, out)
	}
}

func extractCondition(node *ast.Node) string {
	for _, child := range node.Children {
		if child.Type == "condition" || child.Type == "parenthesized_expression" {
			return child.Content
		}
	}
	return ""
}
