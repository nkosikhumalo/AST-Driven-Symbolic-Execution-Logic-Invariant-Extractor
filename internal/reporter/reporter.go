package reporter

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/your-username/axiom/internal/solver"
)

var (
	passStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("10")).Bold(true)
	failStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("9")).Bold(true)
	dimStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("8"))
	separator = dimStyle.Render(string([]rune{45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45}))
)

// Report prints the verification result to stdout.
func Report(result *solver.Result, legacyFile, modernFile string) {
	fmt.Println()

	if !result.Sat {
		fmt.Println(passStyle.Render("EQUIVALENCE VERIFIED"))
		fmt.Println(separator)
		fmt.Printf("  Legacy : %s\n", legacyFile)
		fmt.Printf("  Modern : %s\n", modernFile)
		fmt.Println(separator)
		fmt.Println(dimStyle.Render("  Z3 returned UNSAT — no counterexample exists. Rewrite is 100% safe."))
		fmt.Println()
		return
	}

	fmt.Println(failStyle.Render("EQUIVALENCE VERIFICATION FAILED"))
	fmt.Println(separator)
	fmt.Println("  Reason: Logic discrepancy detected.")
	fmt.Println()
	fmt.Println("  Counterexample (Z3 model):")
	fmt.Println(result.Model)
	fmt.Println(separator)
	fmt.Println()
}
