package leetcode

// 0(n) time complexity, 0(1) space complexity
func PivotIndex(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	leftSum := 0
	for i, num := range nums {
		if leftSum == sum-leftSum-num {
			return i
		}
		leftSum += num
	}

	return -1
}
