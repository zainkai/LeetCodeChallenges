/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func findLeaves(root *TreeNode) [][]int {
    if root == nil {
        return make([][]int, 0)
    }
    
    m := map[int][]int{}
    postOrderTrav(root, m)
    
    leafs := make([][]int, len(m))
    for level, leafArr := range m {
        leafs[level] = leafArr
    }
    
    return leafs
}

func postOrderTrav(root *TreeNode, m map[int][]int) int {    
    level := -1
    if root.Left != nil {
        level = max(level, postOrderTrav(root.Left, m))
    }
    
    if root.Right != nil {
        level = max(level, postOrderTrav(root.Right, m))
    }
    
    m[level+1] = append(m[level+1], root.Val)
    return level+1
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}