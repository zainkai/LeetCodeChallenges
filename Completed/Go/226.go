package LeetCode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)

	return root
}

// iterative
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }
    
    stk := []*TreeNode{}
    tmp := root
    
    for tmp != nil || len(stk) > 0 {
        for tmp != nil {
            // swap
            tmp.Left, tmp.Right = tmp.Right, tmp.Left
            
            stk = append(stk, tmp)
            tmp = tmp.Left
        }
        
        // top
        tmp = stk[len(stk)-1]
        // pop
        stk = stk[:len(stk)-1]
        
        tmp = tmp.Right
    }
    
    return root
}
