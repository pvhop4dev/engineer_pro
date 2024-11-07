package leetcode

func ArrangeCoins(n int) int {
	low := 0
	high := n
	result := 0
	for low <= high {
		mid := low + (high-low)/2
		if mid*(mid+1)/2 == n {
			return mid
		} else if mid*(mid+1)/2 < n {
			result = mid
			low = mid + n/2+1
		} else {
			high = mid - n/2 + 1
		}
	}
	return result
}
