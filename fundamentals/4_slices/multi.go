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
