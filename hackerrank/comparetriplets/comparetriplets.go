package main

func solve(a []int32, b []int32) []int32 {
	var sa, sb int32
	for i, _ := range a {
		if a[i] < b[i] {
			sb++
		} else if a[i] > b[i] {
			sa++
		}
	}

	return []int32{sa, sb}
}
