/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func longestUnivaluePath(root *TreeNode) int {
	if root == nil {
		return 0
	}

	max := 0
	helper(root, root.Val, &max)

	return max
}

func helper(root *TreeNode, pVal int, max *int) int {
	if root == nil {
		return 0
	}

	a := helper(root.Left, root.Val, max)
	b := helper(root.Right, root.Val, max)
	// add left and right trees together
	// if one of the trees doesnt equal the parent value it will return 0
	*max = Max(*max, a+b)

	if root.Val == pVal {
		return Max(a, b) + 1
	}

	return 0
}

func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}