package leetcode

import "container/list"

func ShortestPath(grid [][]int, k int) int {
	rows, cols := len(grid), len(grid[0])
	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	// visited[x][y][k_remain] indicates whether we have visited (x, y) with k_remain obstacles left
	visited := make([][][]bool, rows)
	for i := range visited {
		visited[i] = make([][]bool, cols)
		for j := range visited[i] {
			visited[i][j] = make([]bool, k+1)
		}
	}

	queue := list.New()
	queue.PushBack(State{0, 0, k, 0}) // Start at (0, 0) with k eliminations and 0 steps
	visited[0][0][k] = true

	for queue.Len() > 0 {
		cur := queue.Remove(queue.Front()).(State)

		// If we reach the destination, return the steps
		if cur.x == rows-1 && cur.y == cols-1 {
			return cur.steps
		}

		// Explore neighbors
		for _, dir := range directions {
			nx, ny := cur.x+dir[0], cur.y+dir[1]

			if nx >= 0 && nx < rows && ny >= 0 && ny < cols { // Within bounds
				nk := cur.k
				if grid[nx][ny] == 1 { // Obstacle
					nk--
				}
				if nk >= 0 && !visited[nx][ny][nk] { // Valid state
					visited[nx][ny][nk] = true
					queue.PushBack(State{nx, ny, nk, cur.steps + 1})
				}
			}
		}
	}

	// If we exhaust the queue without reaching the destination, return -1
	return -1
}

type State struct {
	x, y, k, steps int
}
