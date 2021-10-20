/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func closestValue(root *TreeNode, target float64) int {
    res := 0
    diff := float64(1<<63-1)
    
    helper(root, target, &diff, &res)
    
    return res
}

func helper(root *TreeNode, target float64, diff *float64, res *int) {
    if root == nil {
        return
    }
    
    tmpDiff := abs(float64(root.Val) - target)
    if tmpDiff < *diff {
        *diff = tmpDiff
        *res = root.Val
    }
    
    helper(root.Left, target, diff, res)
    helper(root.Right, target, diff, res)
}

func abs(x float64) float64 {
    if x < 0 {
        return -x
    }
    return x
}
