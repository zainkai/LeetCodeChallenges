/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flipEquiv(r1 *TreeNode, r2 *TreeNode) bool {
	if r1 == nil && r2 == nil {
		return true
	} else if r1 == nil || r2 == nil {
		return false
	} else if r1.Val != r2.Val {
		return false
	}

	// check if flipped nodes are equal
	if flipEquiv(r1.Left, r2.Right) && flipEquiv(r1.Right, r2.Left) {
		return true
	}
	return flipEquiv(r1.Left, r2.Left) && flipEquiv(r1.Right, r2.Right)
}