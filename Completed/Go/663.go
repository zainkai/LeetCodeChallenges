/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func checkEqualTree(root *TreeNode) bool {
    rootSum := getSum(root)
    return helper(rootSum, root)
}

func helper(sum int, root *TreeNode) bool {
    if root == nil {
        return false
    } else if root.Left != nil && getSum(root.Left) *2 == sum {
        return true
    } else if root.Right != nil && getSum(root.Right) *2 == sum {
        return true
    }
    return helper(sum, root.Left) || helper(sum, root.Right)
}

func getSum(root *TreeNode) int {
    if root == nil {
        return 0
    }
    return root.Val + getSum(root.Left) + getSum(root.Right)
}
