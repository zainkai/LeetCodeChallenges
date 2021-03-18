import (
  "sort"
)

func highFive(items [][]int) [][]int {
  scoresMap := map[int][]int{}
  for _, item := range items {
    id := item[0]
    score := item[1]
    scoresMap[id] = append(scoresMap[id], score)
  }
  
  result := [][]int{}
  for id, scores := range scoresMap {
    if len(scores) > 5 {
      sort.Ints(scores)
      scores = scores[len(scores)-5:]
    }
    result = append(result, []int{ id, avg(scores) }) 
  }
  
  sort.Slice(result, func(a, b int) bool {
    return result[a][0] < result[b][0]
  })
  return result
}

func avg(arr []int) int {
  sum := 0
  for _, v := range arr {
    sum += v
  }
  return sum /len(arr)
}
