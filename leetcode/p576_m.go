package leetcode

func FindPaths(m int, n int, maxMove int, startRow int, startColumn int) int {
	if maxMove == 0 {
		return 0
	}

	mod := 1_000_000_007
	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	dp := make([][][]int, maxMove+1)
	for i := range dp {
		dp[i] = make([][]int, m)
		for j := range dp[i] {
			dp[i][j] = make([]int, n)
		}
	}

	for move := 1; move <= maxMove; move++ {
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				for _, dir := range directions {
					ni, nj := i+dir[0], j+dir[1]
					if ni < 0 || ni >= m || nj < 0 || nj >= n {
						dp[move][i][j]++
					} else {
						dp[move][i][j] = (dp[move][i][j] + dp[move-1][ni][nj]) % mod
					}
				}
			}
		}
	}

	return dp[maxMove][startRow][startColumn]
}

func FindPaths2(m int, n int, maxMove int, startRow int, startColumn int) int {
	memo := make([][][]int, maxMove+1)
	return findPaths2(m, n, maxMove, startRow, startColumn, memo)
}

func findPaths2(m, n, maxMove, startRow, startColumn int, memo [][][]int) int {
	if maxMove == 0 {
		return 0
	}

	mod := 1_000_000_007
	if memo[maxMove] == nil {
		memo[maxMove] = make([][]int, m)
		for i := range memo[maxMove] {
			memo[maxMove][i] = make([]int, n)
		}
	}

	if memo[maxMove][startRow][startColumn] != 0 {
		return memo[maxMove][startRow][startColumn]
	}

	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	for _, dir := range directions {
		ni, nj := startRow+dir[0], startColumn+dir[1]
		if ni < 0 || ni >= m || nj < 0 || nj >= n {
			memo[maxMove][startRow][startColumn]++
		} else {
			memo[maxMove][startRow][startColumn] = (memo[maxMove][startRow][startColumn] + findPaths2(m, n, maxMove-1, ni, nj, memo)) % mod
		}
	}

	return memo[maxMove][startRow][startColumn]
}
const MOD = 1_000_000_007

func FindPaths3(m int, n int, maxMove int, startRow int, startColumn int) int {
    memo := make(map[[3]int]int)
    
    var dfs func(int, int, int) int
    dfs = func(steps, row, col int) int {
        // Nếu ra khỏi biên, đây là một đường đi hợp lệ
        if row < 0 || row >= m || col < 0 || col >= n {
            return 1
        }
        // Nếu hết bước mà vẫn ở trong lưới thì không hợp lệ
        if steps == 0 {
            return 0
        }
        
        // Kiểm tra xem trạng thái này đã được tính trước đó chưa
        key := [3]int{steps, row, col}
        if val, exists := memo[key]; exists {
            return val
        }
        
        // Di chuyển 4 hướng và cộng dồn kết quả
        result := 0
        directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
        for _, d := range directions {
            result = (result + dfs(steps-1, row+d[0], col+d[1])) % MOD
        }
        
        // Lưu kết quả vào memo
        memo[key] = result
        return result
    }
    
    return dfs(maxMove, startRow, startColumn)
}