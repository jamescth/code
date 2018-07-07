package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}

	M, N, S := m-1, n-1, len(nums1)-1

	for {
		if N < 0 {
			break
		}
		if M < 0 {
			nums1[S] = nums2[N]
			S--
			N--
			continue
		}
		if nums1[M] > nums2[N] {
			nums1[S] = nums1[M]
			M--
			S--
			continue
		}
		nums1[S] = nums2[N]
		N--
		S--
	}
}

/*
func merge(nums1 []int, m int, nums2 []int, n int)  {
    	if n == 0 {
		return
	}

	M, N, sidx := m-1, n-1, len(nums1)-1

	for {
		if N < 0 {
			break
		}
		if M < 0 {
			nums1[sidx] = nums2[N]
			sidx--
			N--
			continue
		}
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
*/
