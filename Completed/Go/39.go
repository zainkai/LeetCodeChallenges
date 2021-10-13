func combinationSum(candidates []int, target int) [][]int {
    res := [][]int{}
    
    helper(0, candidates, target, []int{}, &res)
    
    return res
}

func cpyArr(t []int) []int {
    res := make([]int, len(t))
    copy(res, t)
    
    return res
}

func helper(idx int, candidates []int, target int, path []int, res *[][]int) {
    if target < 0 {
        return
    } else if target == 0 {
        *res = append(*res, cpyArr(path))
        return
    }
    
    for i := idx; i < len(candidates); i++ {
        val := candidates[i]
        target -= val
        path = append(path, val)
        
        helper(i, candidates, target, path, res)
        
        target += val
        path = path[:len(path)-1]
        
    }
}
