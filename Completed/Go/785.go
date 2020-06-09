const (
    RED   = 1
    BLACK = 2
)

func isBipartite(graph [][]int) bool {
    visited := map[int]int{}
    adjList := getAdjList(graph)
    
    for id := range adjList {
        if _, ok := visited[id]; !ok && !helper(id, RED, visited, adjList) {
            return false
        }
    }
    return true
}

func helper(id, color int, visited map[int]int, adjList map[int][]int) bool {
    if col, ok := visited[id]; ok {
        return color == col
    }
    visited[id] = color
    
    nextCol := nextColor(color)
    for _, conn := range adjList[id] {
        res := helper(conn, nextCol, visited, adjList)
        if !res {
            return false
        }
    }
    
    return true
}

func nextColor(col int) int {
    if col == RED {
        return BLACK
    }
    return RED
}

func getAdjList(graph [][]int) map[int][]int {
    list := map[int][]int{}
    
    for i, edges := range graph {
        list[i] = append(list[i], edges...)
    }
    
    return list
}
