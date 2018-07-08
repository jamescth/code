package main

func getRow(rowIndex int) []int {
	ret := []int{1}
	k := rowIndex
	if k == 0 {
		return ret
	}

	i := 1
	for i < k+1 {
		ret = append(ret, 1)
		// fmt.Printf("%d new k %d res %v\n", i, k, ret)
		for j := len(ret) - 2; j > 0; j-- {
			ret[j] += ret[j-1]
		}
		// fmt.Printf("%d new k %d res %v\n", i, k, ret)
		i++
	}
	return ret

}
