package main

func largeGroupPositions(S string) [][]int {
	l := len(S)
	if l < 3 {
		return nil
	}

	var ret [][]int
	i, s, curlen := 0, 0, 1

	for i < l {
		// loop for the same char
		for i+1 < l && S[i] == S[i+1] {
			curlen++
			i++
		}
		//fmt.Printf("%d curlen %d s %d %v\n", i, curlen, s, ret)

		if curlen >= 3 {
			ret = append(ret, []int{s, s + curlen - 1})
			s, curlen = i, 1
		}
		// fmt.Printf("%d curlen %d s %d %v\n", i, curlen, s, ret)
		i++
		s += curlen
		curlen = 1
	}

	return ret

}
