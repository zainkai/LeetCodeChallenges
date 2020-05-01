/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func isValidSequence(root *TreeNode, arr []int) bool {
    return dfs(root, arr, 0)
}

func dfs(root *TreeNode, arr []int, i int) bool {
	if root == nil {
		return false
    } else if i > len(arr)-1 {
        return false
    } else if root.Val != arr[i] {
        return false
    } else if root.Val == arr[i] && i == len(arr) -1 {
        return root.Left == nil && root.Right == nil
    }

    return (
        dfs(root.Left, arr, i+1) ||
        dfs(root.Right, arr, i+1))
}



