// convoluted but technically time O(n) space O(2n) 

func longestConsecutive(nums []int) int {
    m := map[int]int{}
    for _, v := range nums {
        m[v] = v+1
    }
    
    visited := map[int]int{}
    res := 0
    
    for _, v := range nums {
        helper(visited, m, v, &res)
    }
    
    return res
}

func helper(visited, m map[int]int, key int, res *int) int {
    if res, ok := visited[key]; ok {
        return res
    }
    
    if next, ok := m[key]; ok {
        visited[key] = helper(visited, m, next, res) +1
    }
    
    *res = max(*res, visited[key])
    return visited[key]
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
