func getFactors(n int) [][]int {
    res := [][]int{}
    
    helper(n, 2, []int{}, &res)
    
    return res
}

func helper(n, seed int, path []int, res *[][]int) {
    if n <= 1 {
        if len(path) > 1 {
            tmp := make([]int, len(path))
            copy(tmp, path)
            *res = append(*res, tmp)
        }
        return
    }
    
    for i := seed; i <= n; i++{
        if n % i == 0 {
            path = append(path, i)
            helper(n/i, i, path, res)
            path = path[:len(path)-1]
        }
    }
}
