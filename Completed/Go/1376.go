func numOfMinutes(n int, headID int, manager []int, informTime []int) int {
    visited := map[int]int{}
    adjList := managersList(manager)
    res := 0
    
    dfs(headID, 0, adjList, informTime, visited, &res)
    
    return res
}

func managersList(managers []int) map[int][]int {
    adjList := map[int][]int{}
    for subId, manager := range managers {
        adjList[manager] = append(adjList[manager], subId)
    }
    return adjList
}

func dfs(id, time int, managers map[int][]int, informTime []int, visited map[int]int, maxTime *int) {
    if _, ok := visited[id]; ok {
        return
    } else {
        visited[id] = time
        if *maxTime < time {
            *maxTime = time
        }
    }
    
    for _, subId := range managers[id] {
        if _, ok := visited[subId]; !ok {
            dfs(subId, time+informTime[id], managers, informTime, visited, maxTime)
        }
    }
}
