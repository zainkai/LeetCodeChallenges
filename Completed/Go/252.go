import (
    "sort"
)

func canAttendMeetings(intervals [][]int) bool {
    if len(intervals) <= 1 {
        return true
    }
    
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })
    
    sTmp, eTmp := intervals[0][0], intervals[0][1]
    for i := 1; i < len(intervals); i++ {
        if isOverlap([]int{sTmp, eTmp}, intervals[i]) {
            return false
        } else {
            s, e := intervals[i][0], intervals[i][1]
            sTmp, eTmp = s, e
        }
    }
    
    return true
}

func isOverlap(inv1, inv2 []int) bool {
    s1, e1 := inv1[0], inv1[1]
    s2, _ := inv2[0], inv2[1]
    
    if s2 >= s1 && s2 < e1 {
        return true
    }
    return false
}
