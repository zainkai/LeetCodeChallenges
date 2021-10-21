func arrayNesting(nums []int) int {
    res := 0
    visited := map[int]int{}
    
    for i := range nums {
        helper(nums, i, &res, visited)
    }
    
    return res
}

func helper(nums []int, idx int, res *int, visited map[int]int) int {
    if val, ok := visited[idx]; ok {
        return val
    }
    
    visited[idx] = 0
    
    next := nums[idx]
    visited[idx] = helper(nums, next, res, visited) + 1
    
    *res = max(*res, visited[idx])
    
    return visited[idx]
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
