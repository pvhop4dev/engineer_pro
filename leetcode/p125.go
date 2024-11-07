package leetcode

import "strings"

// O(n) time complexity, O(1) space complexity
func IsPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		if !isAlphanumeric(s[i]) {
			i++
			continue
		}
		if !isAlphanumeric(s[j]) {
			j--
			continue
		}
		if !strings.EqualFold(string(s[i]), string(s[j])) {
			return false
		}
		i++
		j--
	}
	return true
}

func isAlphanumeric(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
}
