/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return preOrderDfs(root, -1<<63, 1<<63-1)
}

func preOrderDfs(root *TreeNode, leftBound int, rightBound int) bool {
	if root.Val > leftBound && root.Val < rightBound {
		isValid := true
		if root.Left != nil {
			isValid = isValid && preOrderDfs(root.Left, leftBound, root.Val)
		}
		if root.Right != nil {
			isValid = isValid && preOrderDfs(root.Right, root.Val, rightBound)
		}
		return isValid
	}
	return false

}