/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findLeaves(root *TreeNode) [][]int {
    levels := map[int][]int{}
    
    maxLevel := helper(root, levels)
    res := make([][]int, maxLevel+1)
    
    for level, arr := range levels {
        res[level] = arr
    }
    return res
}

func helper(root *TreeNode, n map[int][]int) int {
    if root == nil {
        return -1
    } else if root.Left == nil && root.Right == nil {
        n[0] = append(n[0], root.Val)
        return 0
    }
    
    level := max(helper(root.Left,n), helper(root.Right, n)) +1
    n[level] = append(n[level], root.Val)
    
    return level
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
