package leetcode

// O(n) time complexity, O(1) space complexity
func IsPerfectSquare(num int) bool {
	for i := 1; i <= num; i++ {
		if i*i == num {
			return true
		}
	}
	return false
}

// O(log n) time complexity, O(1) space complexity
func IsPerfectSquare2(num int) bool {
	low, high := 1, num
	for low <= high {
		mid := low + (high-low)/2
		if mid*mid == num {
			return true
		} else if mid*mid < num {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return false
}
