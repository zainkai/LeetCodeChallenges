func validTree(n int, edges [][]int) bool {
    m := createEdgeMap(edges)
    fmt.Println(m)

	visited :=  map[int]bool{}
    if dfs(0, -1, m, visited) && len(visited) == n {

        return true
    }
    return false
}

// O(V +E)
func dfs(start int, parent int, edgesMap map[int][]int, visited map[int]bool) bool {
	if visited[start] {
		return false
    } else {
        visited[start] = true
    }

    res := true
    for _, v := range edgesMap[start] {
        if parent == v {
            continue
        }
        
        res = res && dfs(v, start, edgesMap, visited)
    }

    return res
}

func createEdgeMap(edges [][]int) map[int][]int {
    m := make(map[int][]int)
	
	for _, edge := range edges {
        start, end := edge[0], edge[1]

        m[start] = append(m[start], end)
        m[end] = append(m[end], start)
    }
	return m
}

