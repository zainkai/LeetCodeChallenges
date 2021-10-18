func isMajorityElement(nums []int, target int) bool {
    count := 0
    for _, num := range nums {
        if num == target {
            count++
        }
        if count > len(nums)/2 {
            return true
        }
    }
    return false
}
