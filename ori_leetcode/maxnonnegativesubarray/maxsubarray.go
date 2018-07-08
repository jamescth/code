package main

func maxset(A []int) []int {
	var ret []int

	l := len(A)
	if l == 0 {
		return nil
	}

	maxV, tempV := -1, 0
	i := 0
	tempidx := 0
	var setTmpidx bool

	for i < l {
		//fmt.Printf("%d A[i] < 0 %v %d\n", i, A[i] < 0, A[i])
		if A[i] < 0 {
			if setTmpidx && tempV > maxV {
				maxV = tempV
				ret = A[tempidx:i]
			}
			setTmpidx = false
			tempV = 0
			i++
			continue
		}
		//fmt.Printf("%d ret %v maxV %d tempV %d tempidx %d\n", i, ret, maxV, tempV, tempidx)

		// positive values
		if setTmpidx == false {
			setTmpidx = true
			tempidx = i
		}
		tempV += A[i]
		// fmt.Printf("%d ret %v maxV %d tempV %d tempidx %d\n", i, ret, maxV, tempV, tempidx)
		i++
	}

	if setTmpidx && tempV > maxV {
		maxV = tempV
		ret = A[tempidx:l]
	}

	if maxV < 0 {
		return nil
	}
	return ret
}
