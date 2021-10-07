func anagramMappings(nums1 []int, nums2 []int) []int {
    m := map[int]int{}
    for i, val := range nums2 {
        m[val] = i
    }
    
    res := make([]int, len(nums1))
    for i, val := range nums1 {
        res[i] = m[val]
    }
    
    return res
}
