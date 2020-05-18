func combine(n int, k int) [][]int {
    results := [][]int{}
    if k > n {
        return results
    }
    
    dfs([]int{}, n, k, 1, &results)
    return results
}

func dfs(buf []int, n, k, idx int, results *[][]int) {
    if k == 0 {
        tmp := make([]int, len(buf))
        copy(tmp, buf)
        *results = append(*results, tmp)
    }
    
    for i := idx; i <= n; i++ {
        buf = append(buf, i)
        
        dfs(buf, n, k-1, i+1, results)
        
        buf = buf[:len(buf)-1]
    }
}