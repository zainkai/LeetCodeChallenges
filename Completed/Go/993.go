/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isCousins(root *TreeNode, x int, y int) bool {    
    xDepth, xParent :=helper(root, x, 0, -1)
    yDepth, yParent := helper(root, y, 0, -1)
    
    return xDepth == yDepth && xParent != yParent
}

func helper(root *TreeNode, target, depth, parent int) (int, int) {
    if root == nil {
        return -1, -1
    } else if root.Val == target {
        return depth, parent
    }
    
    depth++
    parent = root.Val
    
    if res, parent := helper(root.Left, target, depth, parent); res != -1 {
        return res, parent
    }
    return helper(root.Right, target, depth, parent)
}
