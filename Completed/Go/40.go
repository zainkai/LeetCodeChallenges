import (
    "sort"
)

func combinationSum2(candidates []int, target int) [][]int {
    sort.Ints(candidates)
    results := [][]int{}
    
    backtrack(candidates, []int{}, target, 0, &results)
    
    return results
}


func backtrack(candi, buf []int , target, idx int, results *[][]int) {
    if target < 0 {
        return
    } else if target == 0 {
        tmp := make([]int, len(buf))
        copy(tmp, buf)
        *results = append(*results, tmp)
        
        return
    }
    
    for i := idx; i < len(candi); i++ {
        candidate := candi[i]
        buf = append(buf, candidate)
        target -= candidate
        
        backtrack(candi, buf, target, i+1, results)
        
        target += candidate
        buf = buf[:len(buf)-1]
        
        for i < len(candi)-1 && candi[i] == candi[i+1] {
            i++
        }
    }
    
}