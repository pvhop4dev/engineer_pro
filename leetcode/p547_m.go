package leetcode

type UnionFind struct {
	parent []int
	rank   []int
}

func NewUnionFind(size int) *UnionFind {
	parent := make([]int, size)
	rank := make([]int, size)
	for i := 0; i < size; i++ {
		parent[i] = i
	}
	return &UnionFind{parent, rank}
}

func (uf *UnionFind) find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf *UnionFind) union(x, y int) {
	rootX := uf.find(x)
	rootY := uf.find(y)
	if rootX != rootY {
		if uf.rank[rootX] > uf.rank[rootY] {
			uf.parent[rootY] = rootX
		} else if uf.rank[rootX] < uf.rank[rootY] {
			uf.parent[rootX] = rootY
		} else {
			uf.parent[rootY] = rootX
			uf.rank[rootX]++
		}
	}
}

func FindCircleNum(M [][]int) int {
	n := len(M)
	if n == 0 {
		return 0
	}

	uf := NewUnionFind(n)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if M[i][j] == 1 {
				uf.union(i, j)
			}
		}
	}

	circleCount := 0
	for i := 0; i < n; i++ {
		if uf.find(i) == i {
			circleCount++
		}
	}
	return circleCount
}

func FindCircleNum2(M [][]int) int {
	n := len(M)
	if n == 0 {
		return 0
	}

	visited := make([]bool, n)
	count := 0

	var dfs func(int)
	dfs = func(i int) {
		for j := 0; j < n; j++ {
			if M[i][j] == 1 && !visited[j] {
				visited[j] = true
				dfs(j)
			}
		}
	}

	for i := 0; i < n; i++ {
		if !visited[i] {
			count++
			visited[i] = true
			dfs(i)
		}
	}

	return count
}
