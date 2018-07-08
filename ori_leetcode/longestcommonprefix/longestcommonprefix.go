package main

import "fmt"

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
			fmt.Printf("%s\n", strs[i])
			if (i + 1) < L {
				fmt.Printf("%d %d strIdx %d %s %s\n", L, i, strIdx, strs[i], strs[i+1])
				L1, L2 := len(strs[i]), len(strs[i+1])
				if L1 == 0 || L2 == 0 {
					done = true
					break
				}
				if strIdx >= L1 || strIdx >= L2 {
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
