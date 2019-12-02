/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	return helper(root, sum)
}

func helper(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	} else if root.Left == nil && root.Right == nil {
		return sum-root.Val == 0
	}

	sum -= root.Val
	if helper(root.Left, sum) {
		return true
	}
	return helper(root.Right, sum)
}