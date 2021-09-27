// bad
// time O(3N) space O(2N)

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

// GOOD 
// time O(N) space O(height)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }
    
    if p.Val > root.Val && q.Val > root.Val {
        return lowestCommonAncestor(root.Right, p, q)
    } else if p.Val < root.Val && q.Val < root.Val {
        return lowestCommonAncestor(root.Left, p, q)
    }
    return root
}
