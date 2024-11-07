package leetcode

//0(n) time complexity, 0(1) space complexity
func MaxProfit(prices []int) int {
	min := prices[0]
	profit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		} else if prices[i]-min > profit {
			profit = prices[i] - min
		}
	}

	return profit
}

//0(n^2) time complexity, 0(1) space complexity
func MaxProfit2(prices []int) int {
	maxProfit := 0
	for i := 1; i < len(prices); i++ {
		for j := 0; j < i; j++ {
			if prices[i]-prices[j] > maxProfit {
				maxProfit = prices[i] - prices[j]
			}
		}
	}
	return maxProfit
}
