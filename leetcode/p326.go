package leetcode

// O(log3(n)) time complexity, O(1) space complexity
func IsPowerOfThree(n int) bool {
	if n == 0 {
		return false
	}
	for n%3 == 0 {
		n /= 3
	}
	return n == 1
}

// O(log3(n)) time complexity, O(log3(n)) space complexity
func IsPowerOfThree2(n int) bool {
	if n == 1 {
		return true
	}
	if n == 0 || n%3 != 0 {
		return false
	}
	return IsPowerOfThree2(n / 3)
}
