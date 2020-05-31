/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func postorderTraversal(root *TreeNode) []int {
    res := []int{}
    stk := []*TreeNode{root}
    
    for len(stk) > 0 {
        root = stk[len(stk)-1]
        stk = stk[:len(stk)-1]
        
        if root == nil {
            continue
        }
        
        res = append(res, root.Val)
        
        stk = append(stk, root.Left)
        stk = append(stk, root.Right)
        
    }
    
    reverse(res)
    return res
}

func reverse(res []int) {
    i := 0
    j := len(res)-1
    
    for i < j {
        res[i], res[j] = res[j], res[i]
        
        i++
        j--
    }
}