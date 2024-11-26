package leetcode

import "log"

func ValidPath(n int, edges [][]int, source int, destination int) bool {
	// create graph
	graph := make(map[int][]int)
	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}
	log.Printf("graph: %v", graph)

	// create visited
	visited := make(map[int]bool)

	// create queue
	queue := []int{source}

	// bfs
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node == destination {
			return true
		}

		if visited[node] {
			continue
		}

		visited[node] = true

		for _, neighbor := range graph[node] {
			if !visited[neighbor] {
				queue = append(queue, neighbor)
			}
		}
	}

	return false
}
