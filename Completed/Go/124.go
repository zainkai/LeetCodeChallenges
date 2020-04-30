/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func max(vals ...int) int {
	res := vals[0]
	for _, val := range vals {
		if val > res {
			res = val
		}
	}
	return res
}

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	sum := root.Val

	helper(root, &sum)

	return sum
}

func helper(root *TreeNode, sum *int) int {
	if root == nil {
		return 0
	}

	left := helper(root.Left, sum)
	right := helper(root.Right, sum)

	*sum = max(
		*sum,
		left+root.Val,
		right+root.Val,
		left+right+root.Val,
		root.Val,
	)

	return max(root.Val, root.Val+left, root.Val+right)
}