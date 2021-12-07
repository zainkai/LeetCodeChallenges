// BEST


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getDecimalValue(head *ListNode) int {
  res := 0
  for head != nil {
    if head.Val ==  1 {
      res += 1
    }
    res <<= 1
    head = head.Next
  }
  
  res >>= 1
  
  return res
}

//**************************************************************************************//
import (
  "math"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getDecimalValue(head *ListNode) int {
  length := getLength(head) -1
  
  pow := int(math.Pow(float64(2), float64(length)))
  res := 0
  for head != nil {
    if head.Val == 1 {
      res += pow
    }
    
    pow /= 2
    head = head.Next
  }
  
  return res
}

func getLength(head *ListNode) int {
  res := 0
  for head != nil {
    res++
    head = head.Next
  }
  
  return res
}
