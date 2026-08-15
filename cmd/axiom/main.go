package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	legacyFile string
	modernFile string
)

var rootCmd = &cobra.Command{
	Use:   "axiom",
	Short: "Axiom — formal verification engine for legacy rewrites",
	Long:  `Mathematically prove 100% logic equivalence when rewriting legacy code into modern languages.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if legacyFile == "" || modernFile == "" {
			return fmt.Errorf("both --legacy and --modern flags are required")
		}
		fmt.Printf("[AXIOM] Verifying:\n  Legacy: %s\n  Modern: %s\n", legacyFile, modernFile)
		// TODO: wire up AST parser -> symbolic walker -> SMT solver -> reporter
		return nil
	},
}

func init() {
	rootCmd.Flags().StringVar(&legacyFile, "legacy", "", "Path to legacy source file")
	rootCmd.Flags().StringVar(&modernFile, "modern", "", "Path to modern source file")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
