package leetcode

func CanVisitAllRooms(rooms [][]int) bool {
	visited := make([]bool, len(rooms))
	visited[0] = true
	stack := []int{0}

	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		for _, key := range rooms[cur] {
			if !visited[key] {
				visited[key] = true
				stack = append(stack, key)
			}
		}
	}

	for _, v := range visited {
		if !v {
			return false
		}
	}
	return true
}
