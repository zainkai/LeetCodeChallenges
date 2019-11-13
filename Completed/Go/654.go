/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func constructMaximumBinaryTree(nums []int) *TreeNode {
    if len(nums) == 0 { return nil }
    max, mIdx := arrayMax(nums)
    
    return &TreeNode{
        Val: max,
        Left: constructMaximumBinaryTree(nums[:mIdx]),
        Right: constructMaximumBinaryTree(nums[mIdx+1:]),
    }
}

func arrayMax(nums []int) (int, int) {
    max, mIdx := -1, -1
    for i, m := range nums {
        if m > max {
            max = m
            mIdx = i
        }
    }
    return max, mIdx
}