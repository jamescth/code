package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}

	M, N, sidx := m-1, n-1, len(nums1)-1

	for {
		// nothing to merge
		if N < 0 {
			break
		}
		// move everything from nums2
		if M < 0 {
			nums1[sidx] = nums2[N]
			sidx--
			N--
			continue
		}
		// compare size
		if nums1[M] > nums2[N] {
			nums1[sidx] = nums1[M]
			sidx--
			M--
		} else {
			nums1[sidx] = nums2[N]
			sidx--
			N--
		}
	}
}
