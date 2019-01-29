/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func isUnivalTree(root *TreeNode) bool {
	return helper(root, root.Val)
}

func helper(root * TreeNode, v int) bool {
	if root == nil { return true }
	
	return root.Val == v && helper(root.Left, v) && helper(root.Right, v)
}