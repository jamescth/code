package main

import "math/big"

func factorial(
	n *big.Int,
	accumulate *big.Int,
	one *big.Int) *big.Int {
	// cmp returns -1 if n < 1
	// returns 0 if n == 1
	// so here I check n <= 1
	if n.Cmp(one) < 0 || n.Cmp(one) == 0 {
		return accumulate
	}
	accumulate.Mul(accumulate, n)
	n.Sub(n, one)
	return factorial(n, accumulate, one)
}
