package main

func firstBadVersion(vers []bool) int {
	if vers == nil {
		return 0
	}

	L, R := 0, len(vers)-1
	for L <= R {
		M := (L + R) / 2

		// fmt.Printf("%v L %d M %d R %d\n", vers, L, M, R)
		if vers[M] == true {
			R = M - 1
		} else {
			L = M + 1
		}
	}

	return L
}
