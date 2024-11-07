package leetcode

import (
	"math"
	"strings"
	"unicode"
)

// O(n) time complexity, O(1) space complexity
func MyAtoi(s string) int {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return 0
	}

	sign := 1
	index := 0
	if s[index] == '-' {
		sign = -1
		index++
	} else if s[index] == '+' {
		index++
	}

	result := 0
	for ; index < len(s); index++ {
		if !unicode.IsDigit(rune(s[index])) {
			break
		}

		result = result*10 + int(s[index]-'0')

		if result*sign > math.MaxInt32 {
			return math.MaxInt32
		}
		if result*sign < math.MinInt32 {
			return math.MinInt32
		}
	}

	return result * sign
}
