package solver

import (
	"fmt"
	"strings"

	"github.com/your-username/axiom/internal/symbolic"
)

// BuildEquivalenceQuery generates an SMT-LIB2 query that asserts
// the negation of equivalence between two sets of path constraints.
// Z3 finding SAT on this query means a counterexample exists (logic drift).
func BuildEquivalenceQuery(legacy, modern []symbolic.PathConstraint, vars []string) string {
	var b strings.Builder

	// Declare symbolic integer variables
	for _, v := range vars {
		fmt.Fprintf(&b, "(declare-const %s Int)\n", v)
	}

	b.WriteString("\n; Assert negation of equivalence\n")
	b.WriteString("(assert (not (=\n")
	b.WriteString("  ; legacy path\n")
	writeConstraintBlock(&b, legacy)
	b.WriteString("  ; modern path\n")
	writeConstraintBlock(&b, modern)
	b.WriteString(")))\n\n")
	b.WriteString("(check-sat)\n")
	b.WriteString("(get-model)\n")

	return b.String()
}

func writeConstraintBlock(b *strings.Builder, constraints []symbolic.PathConstraint) {
	for _, c := range constraints {
		if c.Condition != "" {
			fmt.Fprintf(b, "  (ite %s true false)\n", c.Condition)
		}
	}
}
