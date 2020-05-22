func removeDuplicates(nums []int) int {
    idx := 0
    m := map[int]bool{}
    for _, v := range nums {
        if !m[v] {
            m[v] = true
            nums[idx] = v
            idx++
        }
    }
    return idx
}
