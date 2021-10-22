func validPath(n int, edges [][]int, start int, end int) bool {
    if start == end {
        return true
    }
    
    adjList := GenAdjList(edges)
    visited := map[int]bool{}
    return helper(start, end, visited, adjList)
}

func helper(curr, target int, visited map[int]bool, adjList map[int][]int) bool {
    if curr == target {
        return true
    } else if _, ok := visited[curr]; ok {
        return false
    }
    visited[curr] = false
    
    for _, next := range adjList[curr] {
        visited[curr] = visited[curr] || helper(next, target, visited, adjList)
    }
    
    return visited[curr]
}

func GenAdjList (edges [][]int) map[int][]int {
    m := map[int][]int{}
    
    for _, edge := range edges {
        a, b := edge[0], edge[1]
        
        m[a] = append(m[a], b)
        m[b] = append(m[b], a)
    }
    
    return m
}
