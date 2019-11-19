/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
	i := 0
	return helper(root, k, &i)

}

func helper(root *TreeNode, k int, iter *int) int {
	if root == nil {
		return 0
	}

	left := helper(root.Left, k, iter)

	*iter += 1
	if *iter == k {
		return root.Val
	}

	right := helper(root.Right, k, iter)

	if left != 0 {
		return left
	} else if right != 0 {
		return right
	}
	return 0
} 