/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {
    sum := 0
    
    path := []int{}
    helper(root, &path, &sum)
    
    return sum
}

func helper(root *TreeNode, path *[]int, sum *int) {
    if root == nil {
        return
    }
    
    *path = append(*path, root.Val)
    
    if root.Left == nil && root.Right == nil {
        *sum += Arrconv(*path)
    }
    
    helper(root.Left, path, sum)
    helper(root.Right, path, sum)
    
    *path = (*path)[:len(*path)-1]
}

func Arrconv(nums []int) int {
    res := 0
    
    mod := 1
    for i := len(nums)-1; i >= 0; i-- {
        res += (mod * nums[i])
        mod *= 10
    }
    
    return res
}
