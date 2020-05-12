func insert(intervals [][]int, newInterval []int) [][]int {
    // intsert newInterval
    interted := false
    for i, interval := range intervals {
        if newInterval[0] < interval[0] {
            temp := append([][]int{}, intervals[:i]...)
            temp = append(temp, newInterval)
            temp = append(temp, intervals[i:]...)
            
            intervals = temp
            
            interted = true
            break
        }
    }
    if len(intervals) == 0 || !interted {
        intervals = append(intervals, newInterval)
    }
    
    fmt.Println(intervals)
    
    
    ans := [][]int{}
    
    s1 := intervals[0][0]
    e1 := intervals[0][1] 
    // merge overlapping intervals
    for i := 1; i < len(intervals); i++ {
        s2, e2 := intervals[i][0], intervals[i][1]
        if s2 <= e1 && e1 < e2 {
            e1 = e2
        } else if e1 < s2 {
            ans = append(ans, []int{s1, e1})
            s1 = s2
            e1 = e2
        } 
    }
    ans = append(ans, []int{s1, e1})
    
    
    return ans
}