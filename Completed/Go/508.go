/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findFrequentTreeSum(root *TreeNode) []int {
    m := map[int]int{}
    
    //O(N)
    helper(root, m)
    
    //O(N)
    maxFreq := 0
    for _, freq := range m {
        if freq > maxFreq {
            maxFreq = freq
        }
    }
    
    //O(N)
    res := []int{}
    for sum, freq := range m {
        if freq == maxFreq {
            res = append(res, sum)
        }
    }
    
    return res
}

func helper(root *TreeNode, m map[int]int) int {
    if root == nil {
        return 0
    }
    
    sum := helper(root.Left, m)
    sum += helper(root.Right, m)
    sum += root.Val
    
    m[sum]++
    
    return sum
}
