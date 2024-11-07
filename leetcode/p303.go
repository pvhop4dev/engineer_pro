package leetcode

type NumArray struct {
	num []int
}

func Constructor(nums []int) NumArray {
	return NumArray{num: nums}
}

// 0(n) time complexity, 0(1) space complexity
func (a *NumArray) SumRange(i int, j int) int {
	sum := 0
	for k := i; k <= j; k++ {
		sum += a.num[k]
	}
	return sum
}
