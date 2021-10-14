func findCircleNum(edges [][]int) int {
	visited := map[int]bool{}
	friendGroups := 0

	for i := 0; i < len(edges); i++ {
		if !visited[i] {
			friendGroups++
			dfs(visited, edges, i)
		}
	}
	return friendGroups
}

func dfs(visited map[int]bool, edges [][]int, student int) {
	if visited[student] {
		return
	}
	visited[student] = true

	for person, isFriend := range edges[student] {
		if isFriend == 1 && !visited[person] {
			dfs(visited, edges, person)
		}
	}
}


// union set
func findCircleNum(isConnected [][]int) int {
    roots := make([]int, len(isConnected))
    for i := 0; i < len(isConnected); i++ {
        roots[i] = i
    }
    
    res := len(isConnected)
    for a := range isConnected {
        for b := range isConnected[a] {
            if isConnected[a][b] == 1 {
                rootA := find(a, roots)
                rootB := find(b, roots)
                
                if rootA != rootB {
                    roots[rootB] = rootA
                    res--
                }
            }
        }
    }
    return res
}

func find(v int, roots []int) int {
    for roots[v] != v {
        v = roots[v]
    }
    return v
}
