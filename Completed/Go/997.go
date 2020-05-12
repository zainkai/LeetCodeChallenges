func findJudge(N int, trust [][]int) int {
	indegree := map[int]int{}
	numTrusts := map[int]int{}
	for _, edge := range trust {
        truster, trusty := edge[0], edge[1]
        indegree[trusty]++
        numTrusts[truster]++
    }

    for i := 1; i <= N; i++ {
        if indegree[i] == N-1 && numTrusts[i] == 0 {
            return i
        }
    }
    return -1
}

