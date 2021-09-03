func maxSubArray(nums []int) int {
  memo := map[int]int{}
  memo[len(nums)-1] = nums[len(nums)-1]
  res := nums[len(nums)-1]
  
  for i:= len(nums)-2; i >= 0; i-- {
    memo[i] = max(nums[i], nums[i]+memo[i+1])
    res = max(res, memo[i])
  }
  
  return res
}

func max(a, b int) int {
  if a > b {
    return a
  }
  return b
}
