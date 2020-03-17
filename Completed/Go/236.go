/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var lca *TreeNode = nil
	dfs(root, p, q, &lca)

	return lca
}

func dfs(root, p, q *TreeNode, lca **TreeNode) bool {
	if root == nil {
		return false
	}

	a := dfs(root.Left, p, q, lca)
	b := dfs(root.Right, p, q, lca)
	c := root == p || root == q

	cnt := 0
	for _, v := range []bool{a, b, c} {
		if v {
			cnt++
		}
	}

	if cnt == 2 {
		*lca = root
	}
	return cnt > 0
}