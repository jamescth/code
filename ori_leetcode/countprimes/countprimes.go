package main

func countPrimes(n int) int {
	flagArray := make([]bool, n)
	for q := 0; q < len(flagArray); q++ {
		flagArray[q] = true
	}
	result := 0
	for i := 2; i < n; i++ {
		if flagArray[i] == true {
			// is Primes
			result++
			// rm it's multiples
			j := 2
			for i*j < n {
				flagArray[i*j] = false
				j++
			}
		}
	}
	return result
	/*
		// too slow
		if n < 2 {
			return 0
		}

		count := n - 2

		for i := 3; i <= n-1; i++ {
			for j := 2; j < i; j++ {
				if i%j == 0 {
					// fmt.Printf("  i %d j %d\n", i, j)
					count--
					break
				}
				// fmt.Printf("i %d j %d\n", i, j)
			}
		}
		return count
	*/
}
