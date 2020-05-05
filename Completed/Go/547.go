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