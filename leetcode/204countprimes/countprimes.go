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

}
