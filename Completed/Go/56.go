import "sort"

func merge(intervals [][]int) [][]int {
  if len(intervals) == 0 || len(intervals) == 1 {
    return intervals
  }
  
  sort.Slice(intervals, func (i, j int) bool {
    return intervals[i][0] < intervals[j][0]
  })
  
  newInternals := make([][]int, 0)
  s1, e1 := intervals[0][0], intervals[0][1]
  for i := 1; i < len(intervals); i++ {
    s2, e2 := intervals[i][0], intervals[i][1]
    if s2 <= e1 && e1 < e2 {
      e1 = e2
    } else if s2 > e1 {
      newInternals = append(newInternals, []int{ s1,e1 })
      s1 = s2
      e1 = e2
    }
  }
  newInternals = append(newInternals, []int{ s1,e1 })
  return newInternals
  
}