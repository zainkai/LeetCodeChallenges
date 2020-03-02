/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 const (
    MIN = -1 << 63
    MAX = 1 << 63-1
)


func bstFromPreorder(preorder []int) *TreeNode {
    i := 0
    return dfs(preorder, &i, MIN, MAX)
}

func dfs(nums []int, i *int, min int, max int) *TreeNode {
    if *i > len(nums)-1 {
        return nil
    } else if val := nums[*i]; val < min || val > max {
        return nil
    }
    
    head := &TreeNode{ nums[*i], nil, nil }
    *i += 1
    
    head.Left = dfs(nums, i, min, head.Val)
    head.Right = dfs(nums, i, head.Val, max)
    
    return head
}