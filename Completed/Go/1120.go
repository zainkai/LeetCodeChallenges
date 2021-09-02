/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maximumAverageSubtree(root *TreeNode) float64 {
    res := float64(0)
    helper(&res, root)
    
    return res
}

func helper(res * float64, root * TreeNode) (int, int) {
    if root == nil {
        return 0, 0
    }
    sum := root.Val
    count := 0
    
    leftSum, leftCount := helper(res, root.Left)
    rightSum, rightCount := helper(res, root.Right)
    
    sum += leftSum + rightSum
    count += leftCount + rightCount
    
    tmp := float64(sum) / float64(count+1)
    *res = max(*res, tmp)
    
    return sum, count+1
}

func max(a, b float64) float64 {
    if a > b {
        return a
    }
    return b
}
