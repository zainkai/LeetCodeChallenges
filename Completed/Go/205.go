import (
  "strings"
  "strconv"
)

func isIsomorphic(s string, t string) bool {
  if len(s) != len(t) {
    return false
  }
  
  return normalize(s) == normalize(t)
}

func normalize(s string) string {
  charNum := 0
  charMap := map[rune]int{}
  
  sb := strings.Builder{}
  for _, r := range s {
    if _, ok := charMap[r]; !ok {
      charMap[r] = charNum
      charNum++
    }
    
    val := strconv.Itoa(charMap[r])
    sb.WriteString(val)
  }
  return sb.String()
}
