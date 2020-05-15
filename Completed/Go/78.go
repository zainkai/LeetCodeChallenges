func subsets(nums []int) [][]int {
    res := [][]int{
        {},
    }
    for i := range nums {
        dfs(nums, []int{}, i, &res)
    }
    
    return res
}

func dfs(nums, frame []int, idx int, res *[][]int) {
    if idx >= len(nums) {
        return
    }
    
    frame = append(frame, nums[idx])
    *res = append(*res, frame)
    
    for i := idx+1; i < len(nums); i++ {
        temp := make([]int, len(frame))
        copy(temp, frame)
        
        dfs(nums, temp, i, res)
    }
}