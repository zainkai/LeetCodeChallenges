func minCostToMoveChips(position []int) int {
  evens := countEvens(position)
  odds := len(position) - evens
  
  if odds > evens {
    return evens
  }
  return odds
}

func countEvens (arr []int) int {
  res := 0
  for _, v := range arr {
    if v % 2 == 0 {
      res++
    }
  }
  
  return res
}
