/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
    depth := 1<< 63-1
    
    helper(root, 1, &depth)
    
    if depth == 1<< 63-1 {
        return 0
    }
    return depth
}

func helper(root *TreeNode, depth int, minDepth *int) {
    if root == nil {
        return
    }
    
    if root.Left == nil && root.Right == nil && depth < *minDepth {
        *minDepth = depth
        return
    }
    
    helper(root.Left, depth+1, minDepth)
    helper(root.Right, depth+1, minDepth)
}
