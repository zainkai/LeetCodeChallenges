/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findTilt(root *TreeNode) int {
  sum := 0
  helper(root, &sum)
  
  return sum
}

func helper(root *TreeNode, sum *int) int {
  if root == nil {
    return 0
  }
  
  res := root.Val
  
  left := 0
  if root.Left != nil {
    left = helper(root.Left, sum)
  }
  
  right := 0
  if root.Right != nil {
    right = helper(root.Right, sum)
  }
  
  *sum += abs(left-right)
  
  return res + left + right
}

func abs(a int) int {
  if a < 0 {
    return -a
  }
  
  return a
}
