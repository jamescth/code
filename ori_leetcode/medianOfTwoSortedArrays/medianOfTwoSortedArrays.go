package main

func findMedianSortedArrays(nums1, nums2 []int) float64 {
	l1, l2 := len(nums1), len(nums2)
	idx1, idx2 := 0, 0

	// handle even later
	end := (l1 + l2) / 2
	end += 1

	n1, n2 := 0, 0
	for {
		var v1, v2 int

		if idx1+idx2 > end {
			break
		}

		var getl1, getl2 bool
		if l1 != 0 && idx1 < l1 {
			getl1 = true
			v1 = nums1[idx1]
		}
		if l2 != 0 && idx2 < l2 {
			getl2 = true
			v2 = nums2[idx2]
		}

		if getl1 && getl2 {

		}

		if getl1 {
			n1 = n2
			n2 = v1
			idx1++
			continue
		}

		// getl2
		n1 = n2
		n2 = v2
		idx2++

	}

	var even bool
	m := (l1 + l2) / 2
	if (l1+l2)%2 == 0 {
		even = true
	} else {
		m++
	}

}
