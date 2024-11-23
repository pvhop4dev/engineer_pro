package leetcode

func CharacterReplacement(s string, k int) int {
	n := len(s)
	begin := 0
	end := 0
	maxCount := 0
	result := 0
	count := make([]int, 26)
	for end < n {
		count[s[end]-'A']++
		maxCount = max(maxCount, count[s[end]-'A'])
		if end-begin+1-maxCount > k {
			count[s[begin]-'A']--
			begin++
		}
		result = max(result, end-begin+1)
		end++
	}
	return result
}

func CharacterReplacement2(s string, k int) int {
	n := len(s)
	begin := 0
	end := n - 1
	maxCount := 0
	cnt := make([]int, 26)
	for begin < n {
		for end > 0 {
			cnt[s[end]-'A']++
			if numChanges(s, begin, end, cnt) > k {
				end--
			} else {
				cnt[s[end]-'A']--
				break
			}
		}
		maxCount = max(maxCount, end-begin+1)
		cnt[s[begin]-'A']--
		begin++
	}
	return maxCount
}

func numChanges(s string, begin int, end int, cnt []int) int {
	max := cnt[0]
	for c := 1; c < 26; c++ {
		if cnt[c] > max {
			max = cnt[c]
		}
	}
	return end - begin + 1 - max
}
