package leetcode
//TODO
// O(log(n)) time complexity, O(1) space complexity
var mapOfSquares = map[int]bool{}

func IsHappy(n int) bool {
	if n == 1 {
		return true
	}
	if _, ok := mapOfSquares[n]; ok {
		return false
	}
	mapOfSquares[n] = true
	return IsHappy(sumOfDigitSquares(n))
}

func sumOfDigitSquares(n int) int {
	sum := 0
	for n > 0 {
		digit := n % 10
		sum += digit * digit
		n /= 10
	}
	return sum
}
