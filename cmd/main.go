package main

import (
	"engineer_pro/leetcode"
	"log"
)

func main() {
	edges := [][]int{{0, 1}, {0, 2}, {3, 5}, {5, 4}, {4, 3}}
	source := 0
	destination := 5
	n := 6
	log.Printf("validPath: %v", leetcode.ValidPath(n, edges, source, destination))
}
