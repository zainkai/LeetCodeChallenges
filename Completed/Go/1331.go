import (
  "sort"
)

func arrayRankTransform(arr []int) []int {
  setRank := map[int]int{}
  for _, v := range arr {
    setRank[v] = 0
  }
  ranks := make([]int, 0, len(setRank))
  for k := range setRank {
    ranks = append(ranks, k)
  }
  sort.Ints(ranks)
  for i, k := range ranks {
    setRank[k] = i+1
  }
  
  for i, v := range arr {
    arr[i] = setRank[v]
  }
  
  return arr
}
