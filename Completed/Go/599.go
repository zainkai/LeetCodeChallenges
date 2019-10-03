import "math"

// Better solution
func findRestaurant(list1 []string, list2 []string) []string {
  m := map[string]int{}
  for i, val := range list1 {
    m[val] = i
  }
  
  sum := math.MaxInt64
  result := []string{}
  for j, val := range list2 {
    if i, ok := m[val]; ok {
      if i+j < sum {
        result = []string{val}
        sum = i+j
      } else if i+j == sum {
        result = append(result, val)
      }
    }
  }
  return result
}

// first solution
import "sort"

func findRestaurant(list1 []string, list2 []string) []string {
  m := map[string]*[2]int{}
  for i, val := range list1 {
    m[val] = &[2]int{i, -1}
  }
  
  inter := []string{}
  for j, val := range list2 {
    if m[val] != nil {
      inter = append(inter, val)
      (*m[val])[1] = j
    }
  }
  
  sort.Slice(inter, func(i,j int) bool{
    keyI := inter[i]
    keyJ := inter[j]
    return m[keyI][0]+ m[keyI][1] < m[keyJ][0] + m[keyJ][1]
  })
  
  sum := m[inter[0]][0] + m[inter[0]][1]
  result := []string{}
  for _, val := range inter {
    if m[val][0] + m[val][1] == sum {
      result = append(result, val)
    }
  }
  
  return result
}