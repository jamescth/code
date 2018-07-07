package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	numStr := len(strs)
	if numStr == 0 {
		return ""
	}
	if numStr == 1 {
		return strs[0]
	}

	maxbyte := minLengthString(strs)
	if maxbyte == 0 {
		return ""
	}
	done := false

	fmt.Printf("%v maxbyte %d\n", strs, maxbyte)
	// compare all bytes until maxbyte
	j := 0
	for !done && j < maxbyte {
		// compare each string
		i := 0
		for i+1 < numStr {
			if j < maxbyte && strs[i][j] != strs[i+1][j] {
				done = true
				j--
				break
			}
			i++
		}
		j++
	}

	return strs[0][:j]
}

func minLengthString(strs []string) int {
	min, i, l := 0, 0, len(strs)
	if l == 0 {
		return 0
	}
	min = len(strs[0])
	for i < l {
		if ret := len(strs[i]); ret < min {
			min = ret
		}
		i++
	}
	return min
}

/*
func longestCommonPrefix(strs []string) string {
	var ret string
	m := minStringL(strs)
	if m == 0 {
		return ret
	}
	fmt.Printf("%v , %d\n", strs, m)

	snum := len(strs)
	if snum == 1 {
		return strs[0]
	}

	j := 0
	for j < m {
		i := 0
		for i+1 < snum {
			fmt.Printf("%v i %d j %d\n", strs, i, j)
			if strs[i][j] != strs[i+1][j] {
				break
			}

			i++
		}
		j++
	}
	ret = strs[0][:j]
	return ret
}

// the min length string
func minStringL(strs []string) int {
	min, i, l := -1, 0, len(strs)

	for i < l {
		ret := len(strs[i])
		fmt.Printf("min %d ret %d\n", min, ret)
		if min < 0 {
			min = ret
		}
		if ret < min {
			min = ret
		}
		i++
	}

	if min < 0 {
		return 0
	}
	return min
}
*/

/*
func longestCommonPrefix(strs []string) string {
	L := len(strs)
	if L == 0 {
		return ""
	}
	if L == 1 {
		return strs[0]
	}

	var ret string
	strIdx := 0
	done := false
	// break when:
	// 1. any string ends
	// 2. not the same char
	// loop strs
	for !done {
		for i := 0; i < L; i++ {
			//fmt.Printf("%s\n", strs[i])
			if (i + 1) < L {
				//fmt.Printf("%d %d strIdx %d %s %s\n", L, i, strIdx, strs[i], strs[i+1])
				L1, L2 := len(strs[i]), len(strs[i+1])
				if L1 == 0 || L2 == 0 {
					done = true
					break
				}

				if strIdx >= len(strs[i]) {
					ret = strs[i][:strIdx]
					done = true
					break
				}
				if strIdx >= len(strs[i+1]) {
					ret = strs[i][:strIdx]
					done = true
					break
				}
				if strs[i][strIdx] != strs[i+1][strIdx] {
					ret = strs[i][:strIdx]
					done = true
					break
				}
			}
		}
		strIdx++
	}
	return ret
}
*/
