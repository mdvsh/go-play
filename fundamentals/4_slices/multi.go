package multi

func Multi(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	prod := 1
	// for i := 0; i < len(nums); i++ {
	// 	prod *= nums[i]
	// }
	for _, num := range nums {
		prod *= num
	}
	return prod
}

func MultiAll(numsToMult ...[]int) (mults []int) {
	maxLen := 0
	for _, nums := range numsToMult {
		if len(nums) > maxLen {
			maxLen = len(nums)
		}
	}
	mults = make([]int, maxLen)
	for i := 0; i < maxLen; i++ {
		mults[i] = 1
	}
	for _, nums := range numsToMult {
		for i := 0; i < len(nums); i++ {
			mults[i] *= nums[i]
		}
	}
	return mults
}
