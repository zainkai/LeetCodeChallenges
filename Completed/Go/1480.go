func runningSum(nums []int) []int {
    sum := 0
    res := []int{}
    
    for _, n := range nums {
        sum += n
        res = append(res, sum)
    }
    return res
}
