package leetcode

func ShortestPathBinaryMatrix(grid [][]int) int {
	n := len(grid)
	if grid[0][0] == 1 || grid[n-1][n-1] == 1 {
		return -1
	}

	// create visited
	visited := make(map[[2]int]bool)

	// create queue
	queue := [][2]int{{0, 0}}

	// bfs
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node[0] == n-1 && node[1] == n-1 {
			return grid[node[0]][node[1]]
		}

		if visited[node] {
			continue
		}

		visited[node] = true

		for _, neighbor := range getNeighbors(node, n) {
			if !visited[neighbor] && grid[neighbor[0]][neighbor[1]] == 0 {
				grid[neighbor[0]][neighbor[1]] = grid[node[0]][node[1]] + 1
				queue = append(queue, neighbor)
			}
		}
	}

	return -1
}

func getNeighbors(node [2]int, n int) [][2]int {
	neighbors := [][2]int{}

	for _, dir := range [][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}} {
		neighbor := [2]int{node[0] + dir[0], node[1] + dir[1]}
		if neighbor[0] >= 0 && neighbor[0] < n && neighbor[1] >= 0 && neighbor[1] < n {
			neighbors = append(neighbors, neighbor)
		}
	}

	return neighbors
}

func ShortestPathBinaryMatrix2(grid [][]int) int {
	n := len(grid)
	xx := []int{-1, -1, -1, 0, 0, 1, 1, 1}
	yy := []int{-1, 0, 1, -1, 1, -1, 0, 1}

	if grid[0][0] == 1 || grid[n-1][n-1] == 1 {
		return -1
	}
	visited := make(map[[2]int]bool)
	queue := [][2]int{{0, 0}}
	grid[0][0] = 1
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node[0] == n-1 && node[1] == n-1 {
			return grid[node[0]][node[1]]
		}
		if visited[node] {
			continue
		}
		visited[node] = true
		for i := 0; i < 8; i++ {
			x := node[0] + xx[i]
			y := node[1] + yy[i]
			if x >= 0 && x < n && y >= 0 && y < n && grid[x][y] == 0 {
				grid[x][y] = grid[node[0]][node[1]] + 1
				queue = append(queue, [2]int{x, y})
			}
		}
	}
	return -1
}
