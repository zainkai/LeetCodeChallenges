type frame struct {
  x, y int
}

func uniquePaths(m int, n int) int {
  memo := map[frame]int{}
  
  return helper(
    frame{0,0},
    frame{m-1,n-1},
    memo,
  )
}

func helper(curr, target frame, memo map[frame]int) int {
  if curr == target {
    return 1
  } else if res, ok := memo[curr]; ok {
    return res
  } else if isNotInBounds(curr, target) {
    return 0
  }
  
  memo[curr] = (helper(frame{curr.x+1, curr.y}, target, memo)+ 
    helper(frame{curr.x, curr.y+1}, target, memo))
  
  return memo[curr]
}

func isNotInBounds(curr, target frame) bool {
  return (curr.x > target.x || curr.y > target.y)
}
