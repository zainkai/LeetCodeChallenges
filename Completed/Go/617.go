/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
    if root1 == nil && root2 == nil {
        return nil
    }
    
    if root1 == nil {
        root1 = &TreeNode{0, nil, nil} 
    }
    if root2 == nil {
        root2 = &TreeNode{0, nil, nil} 
    }
    
    elm := &TreeNode{
        Val: root1.Val+root2.Val,
        Left: mergeTrees(root1.Left, root2.Left),
        Right: mergeTrees(root1.Right, root2.Right),
    }
    return elm
}
