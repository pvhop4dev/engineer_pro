package leetcode

// 0(n) time complexity, 0(1) space complexity
func RemoveElement(nums []int, val int) int {
	cnt := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[cnt] = nums[i]
			cnt++
		}
	}
	return cnt
}
