/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}

	paths := 0
	helper(root, sum, &paths)

	paths += pathSum(root.Left, sum)
	paths += pathSum(root.Right, sum)

	return paths
}

func helper(root *TreeNode, sum int, paths *int) {
	if root == nil {
		return
	}

	sum -= root.Val
	if sum == 0 {
		*paths++
	}

	helper(root.Left, sum, paths)
	helper(root.Right, sum, paths)
}