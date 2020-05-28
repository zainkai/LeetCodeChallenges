const (
    red = 1
    black = 2
)

func possibleBipartition(N int, dislikes [][]int) bool {
    // step 1 create dislike connections
    adjList := getAdjList(dislikes)
    visited := map[int]int{}
    
    // step 2 not all groups know each other
    for i := 1; i <= N; i ++ {
        helper(i, red, visited, adjList)
    }
    
    // step 3 verify groups
    for curr := range visited {
        for _, conn := range adjList[curr] {
            if visited[curr] == visited[conn] {
                return false
            }
        } 
    }
    
    return true
}


func helper(curr int, color int, visited map[int]int, adjList map[int][]int) {
    if _, ok := visited[curr]; ok {
        return
    }
    
    visited[curr] = color
    if color == red {
        color = black
    } else {
        color = red
    }
    
    
    for _, conn := range adjList[curr] {
        helper(conn, color, visited, adjList)
    }
}

func getAdjList(edges [][]int) map[int][]int {
    tmp := map[int][]int{}
    for _, edge := range edges {
        tmp[edge[0]] = append(tmp[edge[0]], edge[1])
        tmp[edge[1]] = append(tmp[edge[1]], edge[0])
    }
    return tmp
}