import (
    "sort"
)

func canAttendMeetings(intervals [][]int) bool {
    if len(intervals) <= 1 {
        return true
    }
    
    sort.Slice(intervals, func(i,j int) bool {
        if intervals[i][0] != intervals[j][0] {
            return intervals[i][0] < intervals[j][0]
        }
        return intervals[i][1] < intervals[j][1]
    })
    
    _, e1 := intervals[0][0], intervals[0][1]
    for i := 1; i < len(intervals); i++ {
        s2, e2 := intervals[i][0], intervals[i][1]
        
        if s2 < e1 {
            return false
        } else if s2 >= e1 {
            e1 = e2
        }
    }
    
    return true
}