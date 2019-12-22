/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	arr := []int{}

	helper(root, &arr)

	return arr
}

func helper(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}

	helper(root.Left, arr)
	*arr = append(*arr, root.Val)
	helper(root.Right, arr)
}