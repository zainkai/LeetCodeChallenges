/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rangeSumBST(root *TreeNode, L int, R int) int {
    sum := 0
    dfs(root, L, R, &sum)
    
    return sum
}

func dfs(root *TreeNode, L int, R int, sum *int) {
    if root == nil {
        return
    } else if root.Val >= L && root.Val <= R {
        *sum += root.Val
    }
    
    dfs(root.Left, L, R, sum)
    dfs(root.Right, L, R, sum)
}
