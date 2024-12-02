package leetcode

import "math"

func LongestIncreasingPath(matrix [][]int) int {
	if len(matrix) == 0 {
		return 0
	}

	rows, cols := len(matrix), len(matrix[0])
	memo := make([][]int, rows)
	for i := range memo {
		memo[i] = make([]int, cols)
	}

	// Các hướng di chuyển: lên, xuống, trái, phải
	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	var dfs func(int, int) int
	dfs = func(row, col int) int {
		// Nếu đã tính toán trước đó, trả về kết quả
		if memo[row][col] != 0 {
			return memo[row][col]
		}

		maxPath := 1
		for _, dir := range directions {
			newRow, newCol := row+dir[0], col+dir[1]
			if newRow >= 0 && newRow < rows && newCol >= 0 && newCol < cols && matrix[newRow][newCol] > matrix[row][col] {
				maxPath = int(math.Max(float64(maxPath), float64(1+dfs(newRow, newCol))))
			}
		}

		memo[row][col] = maxPath
		return maxPath
	}

	maxLength := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			maxLength = int(math.Max(float64(maxLength), float64(dfs(i, j))))
		}
	}

	return maxLength
}

func LongestIncreasingPath2(matrix [][]int) int {
	if len(matrix) == 0 {
		return 0
	}

	rows, cols := len(matrix), len(matrix[0])
	inDegree := make([][]int, rows)
	for i := range inDegree {
		inDegree[i] = make([]int, cols)
	}

	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	// Tính toán "in-degree" cho mỗi ô
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			for _, d := range directions {
				ni, nj := i+d[0], j+d[1]
				if ni >= 0 && ni < rows && nj >= 0 && nj < cols && matrix[ni][nj] > matrix[i][j] {
					inDegree[ni][nj]++
				}
			}
		}
	}

	// Bắt đầu từ các ô có "in-degree" bằng 0
	queue := make([][2]int, 0)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if inDegree[i][j] == 0 {
				queue = append(queue, [2]int{i, j})
			}
		}
	}

	// BFS để tìm đường đi dài nhất
	pathLength := 0
	for len(queue) > 0 {
		pathLength++
		nextQueue := make([][2]int, 0)
		for _, cell := range queue {
			x, y := cell[0], cell[1]
			for _, d := range directions {
				ni, nj := x+d[0], y+d[1]
				if ni >= 0 && ni < rows && nj >= 0 && nj < cols && matrix[ni][nj] > matrix[x][y] {
					inDegree[ni][nj]--
					if inDegree[ni][nj] == 0 {
						nextQueue = append(nextQueue, [2]int{ni, nj})
					}
				}
			}
		}
		queue = nextQueue
	}

	return pathLength
}
