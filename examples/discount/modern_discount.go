// This file is an example for Axiom verification — not part of the Go build.
//go:build ignore

package main

func CalculateDiscount(amount int, tier int) int {
	if amount >= 1000 { // BUG: >= instead of >
		if tier == 2 {
			return amount - 100
		}
		return amount - 50
	}
	return amount
}
