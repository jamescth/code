package main

func reverse(nums []int) {
	i, j := 0, len(nums)-1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}

func rotate(nums []int, k int) {
	l := len(nums)
	if l == 1 || k == 0 {
		return
	}

	k = k % l
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

/*
func rotate(nums []int, k int)  {
    	l := len(nums)
    if l == 1 || k==0 {
        return
    }

    k = k%l
	r := l - k
    fmt.Println(r)
    tmp := append(nums[r:l], nums[:r]...)

    for i:=0 ; i<l ; i++ {
        nums[i] = tmp[i]
    }
}
*/
