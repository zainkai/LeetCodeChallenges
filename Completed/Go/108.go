/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
    if len(nums) == 0 {
        return nil
    } else if len(nums) == 1 {
        return &TreeNode{
            nums[0],
            nil,
            nil,
        }
    }
    
    mid := len(nums)/2
    return &TreeNode{
        nums[mid],
        sortedArrayToBST(nums[:mid]),
        sortedArrayToBST(nums[mid+1:]),
    }
}
