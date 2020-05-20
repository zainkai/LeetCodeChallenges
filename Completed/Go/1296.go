func isPossibleDivide(nums []int, k int) bool {
    m := map[int]int{}
    for _, val := range nums {
        m[val]++
    }
    

    for len(m) != 0 {
        key := getMinKey(m)
        if !helper(key, k, m) {
            return false
        }
    }

    return true
}

func helper(num, k int, m map[int]int) bool {
    if _, ok := m[num]; !ok {
        return false
    }
    
    m[num]--
    if m[num] == 0 {
        delete(m, num)
    }
    
    if k == 1 {
        return true
    }
    return helper(num+1, k-1, m)
}

func getMinKey(m map[int]int) int {
    min := 1 << 63-1
    for key := range m {
        if key < min {
            min = key
        }
    }
    return min
}