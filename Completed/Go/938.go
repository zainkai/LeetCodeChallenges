/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rangeSumBST(root *TreeNode, low int, high int) int {
  if root == nil {
    return 0
  }
  
  res := 0
  if root.Val >= low && root.Val <= high {
    res += root.Val
  }
  
  res += rangeSumBST(root.Left, low, high)
  res += rangeSumBST(root.Right, low, high)
  
  return res
}
