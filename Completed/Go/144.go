/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func preorderTraversal(root *TreeNode) []int {
    res := []int{}
    
    stk := []*TreeNode{root}
    for len(stk) > 0 {
        root = stk[len(stk)-1]
        stk = stk[:len(stk)-1]
        
        if root == nil {
            continue
        }
        
        res = append(res, root.Val)
        
        stk = append(stk, root.Right)
        stk = append(stk, root.Left)
    }
    
    return res
}