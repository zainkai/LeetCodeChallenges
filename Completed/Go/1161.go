/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxLevelSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	m := make(map[int]int)
	dfs(root, 1, m)

	level := 0
	for l := range m {
		if m[l] > m[level] {
			level = l
		}
	}

	return level
}

func dfs(root *TreeNode, level int, m map[int]int) {
	if root == nil {
		return
	}

	dfs(root.Left, level+1, m)
	m[level] += root.Val
	dfs(root.Right, level+1, m)
}