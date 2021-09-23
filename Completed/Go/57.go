func insert(intervals [][]int, newInterval []int) [][]int {
    //intervals = append(intervals, newInterval)
    //sort.Slice(intervals, func(i,j int) bool {
    //    return intervals[i][0] < intervals[j][0]
    //})
    
    // insert newInterval
    inserted := false
    for i, interval := range intervals {
        if newInterval[0] <= interval[0] {
            intervals = append(intervals[:i], append([][]int{newInterval}, intervals[i:]...)...)
            inserted = true
            break
        }
    }
    if !inserted {
        intervals = append(intervals, newInterval)
    }
    
    // merge intervals
    res := [][]int{}
    start := intervals[0][0]
    end := intervals[0][1]
    
    for i := 1; i < len(intervals); i++ {
        tStart := intervals[i][0]
        tEnd := intervals[i][1]
        if tStart <= end && end < tEnd {
            end = tEnd
        } else if tStart > end {
            res = append(res, []int{start, end})
            start = tStart
            end = tEnd
        }
    }
    res = append(res, []int{start, end})
    
    
    return res
}
