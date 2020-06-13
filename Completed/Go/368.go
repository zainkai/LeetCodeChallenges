import (
    "sort"
)

func largestDivisibleSubset(nums []int) []int {
    sort.Ints(nums)
    
    longestPath := []int{}
    
    visited := map[int][]int{}
    for i := len(nums)-1; i >= 0; i-- {                
        tmpPath := helper(i, nums, visited)
        
        if len(tmpPath) > len(longestPath) {
            longestPath = tmpPath
        }
    }
    return longestPath
}

func helper(idx int, nums []int, visited map[int][]int) []int {
    root := nums[idx]
    if path, ok := visited[root]; ok {
        return path
    }
    visited[root] = []int{}
    
    for i := idx-1; i >= 0; i-- {
        num := nums[i]
        if root % num == 0 {
            path := helper(i, nums, visited)
            
            if len(path) > len(visited[root]) {
                visited[root] = path
            }
        }
    }
    
    visited[root] = append(copyArr(visited[root]), root)
    return visited[root]
}

func copyArr(arr []int) []int {
    res := make([]int, len(arr))
    copy(res, arr)
    
    return res
}
