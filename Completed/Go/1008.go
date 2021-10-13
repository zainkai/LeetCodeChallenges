/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func bstFromPreorder(preorder []int) *TreeNode {
    return helper(&preorder, -1<<63, 1<<63-1)
}

func helper(preorder *[]int, min, max int) *TreeNode {
    if len(*preorder) == 0 {
        return nil
    } 
    val := (*preorder)[0]
    if val < min || val > max {
        return nil
    } else {
        *preorder = (*preorder)[1:]
    }
    
    return &TreeNode{
        Val: val,
        Left: helper(preorder, min, val),
        Right: helper(preorder, val, max),
    }
}
