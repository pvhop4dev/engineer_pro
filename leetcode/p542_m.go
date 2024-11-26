package leetcode

func UpdateMatrix(mat [][]int) [][]int {
	n := len(mat)
	xx := []int{-1, 0, 1, 0}
	yy := []int{0, -1, 0, 1}
	visted := make(map[[2]int]bool)
	queue := [][2]int{}
	for i := 0; i < n; i++ {
		for j := 0; j < len(mat[i]); j++ {
			if mat[i][j] == 0 {
				queue = append(queue, [2]int{i, j})
				visted[[2]int{i, j}] = true
			}
		}
	}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		for i := 0; i < 4; i++ {
			x := node[0] + xx[i]
			y := node[1] + yy[i]
			if x >= 0 && x < n && y >= 0 && y < len(mat[x]) && !visted[[2]int{x, y}] {
				mat[x][y] = mat[node[0]][node[1]] + 1
				queue = append(queue, [2]int{x, y})
				visted[[2]int{x, y}] = true
			}
		}
	}
	return mat
}

func UpdateMatrix2(mat [][]int) [][]int {
	n := len(mat)
	xx := []int{-1, 0, 1, 0}
	yy := []int{0, -1, 0, 1}
	queue := [][2]int{}
	for i := 0; i < n; i++ {
		for j := 0; j < len(mat[i]); j++ {
			if mat[i][j] == 0 {
				queue = append(queue, [2]int{i, j})
			} else {
				mat[i][j] = -1
			}
		}
	}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		for i := 0; i < 4; i++ {
			x := node[0] + xx[i]
			y := node[1] + yy[i]
			if x >= 0 && x < n && y >= 0 && y < len(mat[x]) && mat[x][y] == -1 {
				mat[x][y] = mat[node[0]][node[1]] + 1
				queue = append(queue, [2]int{x, y})
			}
		}
	}
	return mat
}
