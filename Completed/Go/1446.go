func maxPower(s string) int {
  if len(s) <= 1 {
    return len(s)
  }
  
  tmp, result := 1,0
  for i := 1; i < len(s); i++ {
    prev, curr := s[i-1], s[i]
    if prev == curr {
      tmp++
    } else {
      tmp = 1
    }
    
    if tmp > result {
      result = tmp
    }
  }
  return result
}
