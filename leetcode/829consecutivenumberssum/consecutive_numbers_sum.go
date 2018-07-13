package main

func consecutiveNumbersSum(N int) int {
	res := 0
	for i := 1; i*i < 2*N; i++ {
		// fmt.Printf("i %d \n", i)
		if i%2 == 1 {
			if N%i == 0 {
				// fmt.Printf("i %d n %d \n", i, N)
				res += 1
			}
		} else {
			q := int(i / 2)
			if N%q == 0 && (N/q)%2 == 1 {
				// fmt.Printf("i %d n %d q %d\n", i, N, q)
				res += 1
			}
		}
	}
	return res
}
