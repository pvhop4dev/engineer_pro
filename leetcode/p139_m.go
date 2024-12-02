package leetcode

func WorkBreak(s string, wordDict []string) bool {
	wordSet := make(map[string]bool)
	for _, word := range wordDict {
		wordSet[word] = true
	}

	n := len(s)
	dp := make([]bool, n+1)
	dp[0] = true

	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordSet[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}

	return dp[n]
}

func WordBreak2(s string, wordDict []string) bool {
	// Chuyển wordDict thành map để tra cứu nhanh
	wordSet := make(map[string]bool)
	for _, word := range wordDict {
		wordSet[word] = true
	}

	// Memoization map lưu trữ trạng thái của mỗi vị trí trong chuỗi
	memo := make(map[int]bool)

	var backtrack func(start int) bool
	backtrack = func(start int) bool {
		// Nếu duyệt hết chuỗi, trả về true
		if start == len(s) {
			return true
		}
		// Nếu đã tính toán kết quả tại vị trí này trước đó, trả về kết quả lưu trữ
		if val, exists := memo[start]; exists {
			return val
		}
		// Thử tất cả các cách chia chuỗi từ vị trí hiện tại
		for end := start + 1; end <= len(s); end++ {
			if wordSet[s[start:end]] && backtrack(end) {
				memo[start] = true
				return true
			}
		}
		// Nếu không tìm được cách phân chia hợp lệ, lưu kết quả và trả về false
		memo[start] = false
		return false
	}

	return backtrack(0)
}
