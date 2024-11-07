package leetcode

import "fmt"

// O(n) time complexity, O(n) space complexity
func FindDisappearedNumbers(nums []int) []int {
	m := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = true
	}
	res := make([]int, 0)
	for i := 1; i <= len(nums); i++ {
		if _, ok := m[i]; !ok {
			res = append(res, i)
		}
	}
	return res
}

// O(n) time complexity, O(1) space complexity
func FindDisappearedNumbers2(nums []int) []int {
	for i := 1; i <= len(nums); i++ {
		if nums[i-1] != i {
			if nums[i-1] != nums[nums[i-1]-1] {
				nums[i-1], nums[nums[i-1]-1] = nums[nums[i-1]-1], nums[i-1]
				i--
			}
		}
	}
	fmt.Println(nums)
	res := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			res = append(res, i+1)
		}
	}
	return res
}
