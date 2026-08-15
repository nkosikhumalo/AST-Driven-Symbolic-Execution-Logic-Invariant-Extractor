package solver

import (
	"fmt"
	"os/exec"
	"strings"
)

// Result holds the Z3 solver response.
type Result struct {
	Sat       bool   // true = counterexample found (FAIL), false = UNSAT (PASS)
	Model     string // raw Z3 model output
	RawOutput string
}

// RunZ3 writes the SMT-LIB2 query to Z3 via stdin and parses the response.
func RunZ3(query string) (*Result, error) {
	cmd := exec.Command("z3", "-in")
	cmd.Stdin = strings.NewReader(query)

	out, err := cmd.Output()
	if err != nil {
		// z3 exits non-zero on SAT with a model; check output before failing
		if len(out) == 0 {
			return nil, fmt.Errorf("z3 execution failed: %w", err)
		}
	}

	raw := string(out)
	result := &Result{RawOutput: raw}

	lines := strings.SplitN(strings.TrimSpace(raw), "\n", 2)
	if len(lines) == 0 {
		return nil, fmt.Errorf("empty z3 output")
	}

	switch strings.TrimSpace(lines[0]) {
	case "unsat":
		result.Sat = false
	case "sat":
		result.Sat = true
		if len(lines) > 1 {
			result.Model = lines[1]
		}
	default:
		return nil, fmt.Errorf("unexpected z3 response: %s", lines[0])
	}

	return result, nil
}
