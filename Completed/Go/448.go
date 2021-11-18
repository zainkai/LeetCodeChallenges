func findDisappearedNumbers(nums []int) []int {
    for _, v := range nums {
        idx := abs(v)-1
        if nums[idx] < 0 {
            continue
        } else {
            nums[idx] = -nums[idx]
        }
    }
    
    res := []int{}
    for i, v := range nums {
        if v > 0 {
            res = append(res, i+1)
        }
    }
    
    return res
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
