func combinationSum(candidates []int, target int) [][]int {
    results := [][]int{}
    
    backtrack(candidates, []int{}, target, 0, &results)
    
    return results
}

func backtrack(candidates, buf []int, target, idx int, results *[][]int) {
    if target == 0 {
        tmp := make([]int, len(buf))
        copy(tmp, buf)
        
        *results = append(*results, tmp)
    } else if target < 0 {
        return
    }
    
    for i := idx; i < len(candidates); i++ {
        buf = append(buf, candidates[i])
        t := target - candidates[i]
        backtrack(candidates, buf, t, i, results)
        
        buf = buf[:len(buf)-1]
    }
    
}