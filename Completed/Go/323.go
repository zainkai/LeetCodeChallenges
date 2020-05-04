func countComponents(n int, edges [][]int) int {
	visited := map[int]bool{}
	numClusters := 0

	vMap := GetVertexMap(edges)

	for i := 0; i < n; i++ {
		if !visited[i] {
			dfs(i, visited, vMap)
			numClusters++
		}
	}

	return numClusters
}

func dfs(v int, visited map[int]bool, conns map[int][]int) {
	if visited[v] {
		return
	}
	visited[v] = true
	for _, n := range conns[v] {
		dfs(n, visited, conns)
	}
}

func GetVertexMap(edges [][]int) map[int][]int {
	vertexMap := make(map[int][]int)
	for _, edge := range edges {
		s, e := edge[0], edge[1]

		vertexMap[s] = append(vertexMap[s], e)
		vertexMap[e] = append(vertexMap[e], s)
	}
	return vertexMap
}