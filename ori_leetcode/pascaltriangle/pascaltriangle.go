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
	/*
		res := make([]int, 1, rowIndex+1)
		res[0] = 1
		if rowIndex == 0 {
			return res
		}

		fmt.Printf("new k %d res %v\n", rowIndex, res)

		for i := 0; i < rowIndex; i++ {
			res = append(res, 1)
			fmt.Printf("%d res %v\n", i, res)
			for j := len(res) - 2; j > 0; j-- {
				res[j] += res[j-1]
			}
			fmt.Printf("res %v\n", res)
		}

		return res
	*/
}
