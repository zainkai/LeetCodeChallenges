/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	return helperFunc(root, 0)
}

func helperFunc(root *TreeNode, currDepth int) int {
	if root == nil {
		return currDepth
	}

	left := helperFunc(root.Left, currDepth+1)
	right := helperFunc(root.Right, currDepth+1)

	if left < right {
		return right
	} else {
		return left
	}

}