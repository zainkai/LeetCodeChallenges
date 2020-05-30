func findOrder(numCourses int, prerequisites [][]int) []int {
    adjList, indegree := getMaps(prerequisites)
    
    q := []int{}
    for i := 0; i < numCourses; i++ {
        if indegree[i] == 0 {
            q = append(q, i)
        }
    }
    
    path := []int{}
    for len(q) > 0 {
        course := q[0]
        path = append(path, course)
        q = q[1:]
        
        for _, conn := range adjList[course] {
            indegree[conn]--
            if deg, ok := indegree[conn]; ok && deg == 0 {
                delete(indegree, conn)
                q = append(q, conn)
            }
        }
    }
    
    if len(indegree) == 0{
        return path
    }
    return []int{}
}

func getMaps(edges [][]int) (map[int][]int, map[int]int) {
    adjList := map[int][]int{}
    indegree := map[int]int{}
    
    for _, edge := range edges {
        a, b := edge[0], edge[1]
        indegree[a]++

        adjList[b] = append(adjList[b], a)
    }
    
    return adjList, indegree
}