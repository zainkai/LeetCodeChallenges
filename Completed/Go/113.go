/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func pathSum(root *TreeNode, sum int) [][]int {
    results := [][]int{}
    path := []int{}
    
    helper(root, sum, path, &results)
    
    return results
}

func helper(root *TreeNode, sum int, path []int, result *[][]int) {
    if root == nil {
        return
    } 
    
    sum -= root.Val
    if sum == 0 && root.Left == nil && root.Right == nil {
        tmp := make([]int, len(path))
        copy(tmp, path)
        tmp = append(tmp, root.Val)
        
        *result = append(*result, tmp)
        return
    }
    
    path = append(path, root.Val)
    
    helper(root.Left, sum, path, result)
    helper(root.Right, sum, path, result)
    
    path = path[:len(path)-1]
}