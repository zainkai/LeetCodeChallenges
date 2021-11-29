func permute(nums []int) [][]int {
    res := [][]int{}
    helper(nums, 0, &res)
    
    return res
}

func helper(arr []int, idx int, res *[][]int) {
    if idx == len(arr) {
        *res = append(*res, deepCopy(arr))
    }
    
    for i := idx; i < len(arr); i++ {
        arr[idx], arr[i] = arr[i], arr[idx]
        helper(arr, idx+1, res)
        arr[idx], arr[i] = arr[i], arr[idx]
    }
} 

func deepCopy(arr []int) []int {
    tmp := make([]int, len(arr))
    copy(tmp, arr)
    
    return tmp
}
