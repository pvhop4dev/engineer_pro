package leetcode

func LengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	begin := 0
	end := 0
	result := 0
	m := make(map[byte]int)
	for end < len(s) {
		if _, ok := m[s[end]]; ok {
			begin = max(begin, m[s[end]]+1)
		}
		m[s[end]] = end
		result = max(result, end-begin+1)
		end++
	}
	return result
}

func LengthOfLongestSubstring2(s string) int {
	if len(s) == 0 {
		return 0
	}

	begin := 0
	end := 0
	result := 0
	m := make(map[byte]int)
	for end < len(s) {
		m[s[end]]++
		for m[s[end]] > 1 {
			m[s[begin]]--
			begin++
		}
		result = max(result, end-begin+1)
		end++
	}
	return result
}

func LengthOfLongestSubstring3(s string) int {
	n := len(s)
	begin := 0
	end := -1
	best := 0
	mapChar := make(map[byte]bool)
	for begin < n {
		for end < n-1 && !mapChar[s[end+1]] {
			end++
			mapChar[s[end]] = true
		}
		best = max(best, end-begin+1)

		
	}
	return best
}