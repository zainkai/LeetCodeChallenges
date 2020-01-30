/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deepestLeavesSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	depth, sum := 0, 0

	helper(root, 0, &depth, &sum)

	return sum
}

func helper(root *TreeNode, currDepth int, maxDepth *int, sum *int) {
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil {
		if currDepth == *maxDepth {
			*sum += root.Val
		} else if currDepth > *maxDepth {
			*maxDepth = currDepth
			*sum = root.Val
		}

		return
	}

	helper(root.Left, currDepth+1, maxDepth, sum)
	helper(root.Right, currDepth+1, maxDepth, sum)
}