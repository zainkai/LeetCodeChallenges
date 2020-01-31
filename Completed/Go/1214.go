/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func twoSumBSTs(root1 *TreeNode, root2 *TreeNode, t int) bool {
	m := make(map[int]bool)
	getComplements(root1, t, m)

	return canGetSum(root2, t, m)
}

func canGetSum(root *TreeNode, t int, m map[int]bool) bool {
	if root == nil {
		return false
	} else if m[root.Val] {
		return true
	}

	return canGetSum(root.Left, t, m) || canGetSum(root.Right, t, m)
}

func getComplements(root *TreeNode, t int, m map[int]bool) {
	if root == nil {
		return
	}

	comp := t - root.Val
	m[comp] = true

	getComplements(root.Left, t, m)
	getComplements(root.Right, t, m)
}