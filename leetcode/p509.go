package leetcode

// 0(2^n) time complexity, 0(n) space complexity
func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

// 0(n) time complexity, 0(n) space complexity
func Fibonacci2(n int) int {
	if n <= 1 {
		return n
	}
	f := make([]int, n+1)
	f[0] = 0
	f[1] = 1
	for i := 2; i <= n; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f[n]
}

var memo = map[int]int{}

// 0(n) time complexity, 0(n) space complexity
func Fibonacci3(n int) int {
	if n <= 1 {
		return n
	}
	if _, ok := memo[n]; !ok {
		memo[n] = Fibonacci3(n-1) + Fibonacci3(n-2)
	}
	return memo[n]
}
