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



///********* new
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    pMap := map[*TreeNode]bool{}
    qMap := map[*TreeNode]bool{}

    hasElm(root, p, pMap)
    hasElm(root, q, qMap)

    return helper(root, pMap, qMap)
}

func hasElm(root, target *TreeNode, visited map[*TreeNode]bool) bool {
    if root == nil {
        return false
    } else if root == target {
        visited[root] = true
        return true
    }

    visited[root] = hasElm(root.Left, target, visited) || hasElm(root.Right, target, visited)
    return visited[root]
}

func helper(root *TreeNode, pMap, qMap map[*TreeNode]bool) *TreeNode {
    if root == nil {
        return nil
    }

    if r := helper(root.Left, pMap, qMap); r != nil {
        return r
    } else if r := helper(root.Right, pMap, qMap); r != nil {
        return r
    }

    if pMap[root] && qMap[root] {
        return root
    }
    return nil
}
