func allPathsSourceTarget(graph [][]int) [][]int {
    visited := map[int]bool{}
    path := []int{0}
    res := [][]int{}
    
    helper(graph, &res, path, visited)
    
    return res
}

func helper(graph [][]int, res *[][]int, path []int, visited map[int]bool) {
    curr := path[len(path)-1]
    if _, ok := visited[curr]; ok {
        return
    } else if curr == len(graph)-1 {
        tmp := make([]int, len(path))
        copy(tmp, path)
        
        *res = append(*res, tmp)
    }
    
    visited[curr] = true
    
    for _, next := range graph[curr] {
        path = append(path, next)
        helper(graph, res, path, visited)
        path = path[:len(path)-1]
    }
    
    delete(visited, curr)
}
