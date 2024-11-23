package leetcode

func maxNumEdgesToRemove(n int, edges [][]int) int {
	par := []int{}
	for i := 0; i < n; i++ {
		par[i] = i
	}
	saveCommon := findSave(n, edges, par, 3)
	saveAlice := findSave(n, edges, par, 1)
	if saveAlice == -1 {
		return -1
	}
	saveBob := findSave(n, edges, par, 2)
	if saveBob == -1 {
		return -1
	}
	return saveCommon + saveAlice + saveBob
}

func findSave(n int, edges [][]int, par []int, edgeType int) int {
	save := 0
	for _, edge := range edges {
		if edge[0] != edgeType {
			continue
		}
		p1 := find(edge[1]-1, par)
		p2 := find(edge[2]-1, par)
		if p1 == p2 {
			save++
			continue
		} else if p1 < p2 {
			par[p2] = p1
		} else {
			par[p1] = p2
		}
	}
	if edgeType == 3 {
		return save
	}
	for i := 0; i < n; i++ {
		p := find(i, par)
		if p != 0 {
			return -1
		}
	}
	return save
}

func find(x int, par []int) int {
	if par[x] != x {
		par[x] = find(par[x], par)
	}
	return par[x]
}
