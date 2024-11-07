package leetcode

import "strings"

// O(n*m) where n is the number of strings and m is the length of the shortest string
func LongestCommonPrefix(strs []string) string {

	prefix := strings.Builder{}

	for i := 0; i < len(strs[0]); i++ {
		ok := true
		c := strs[0][i]
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) || strs[j][i] != c {
				ok = false
				break
			}
		}
		if !ok {
			break
		}
		prefix.WriteByte(c)
	}
	return prefix.String()
}
