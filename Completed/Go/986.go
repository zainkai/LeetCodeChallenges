import (
    "sort"
)

func sortIntervals(intervals [][]int) {
    sort.Slice(intervals, func(i,j int) bool {
        if intervals[i][0] != intervals[j][0] {
            return intervals[i][0] < intervals[j][0]
        }
        return intervals[i][1] < intervals[j][1]
    })
}


func intervalIntersection(A [][]int, B [][]int) [][]int {
    //sortIntervals(A)
    //sortIntervals(B)
    
    aIdx, bIdx := 0,0
    intersections := [][]int{}
    
    for aIdx < len(A) && bIdx < len(B) {
        s1, e1 := A[aIdx][0], A[aIdx][1]
        s2, e2 := B[bIdx][0], B[bIdx][1]
        
        s := max(s1, s2)
        e := min(e1, e2)
        if s <= e {
            intersections = append(intersections, []int{s, e})
        }
        
        if e1 < e2 {
            aIdx++
        } else {
            bIdx++
        }
    }
    
    return intersections
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}