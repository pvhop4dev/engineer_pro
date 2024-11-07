package leetcode

// 0(n) time complexity, 0(n) space complexity
func ProductExceptSelf(nums []int) []int {
	mapZeroIndex := make(map[int]bool, 0)
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			mapZeroIndex[i] = true
		}
	}
	zeroCount := len(mapZeroIndex)
	if zeroCount > 1 {
		return make([]int, len(nums))
	}
	if zeroCount == 1 {
		product := 1
		for i := 0; i < len(nums); i++ {
			if _, ok := mapZeroIndex[i]; !ok {
				product *= nums[i]
			}
		}
		for i := 0; i < len(nums); i++ {
			if _, ok := mapZeroIndex[i]; ok {
				nums[i] = product
			} else {
				nums[i] = 0
			}
		}
	}
	if zeroCount == 0 {
		product := 1
		for i := 0; i < len(nums); i++ {
			product *= nums[i]
		}
		for i := 0; i < len(nums); i++ {
			nums[i] = product / nums[i]
		}
	}
	return nums
}
