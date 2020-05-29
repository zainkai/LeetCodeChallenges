/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func increasingBST(root *TreeNode) *TreeNode {
    nodeList := []*TreeNode{}
    helper(root, &nodeList)
    
    for i, n := range nodeList {
        n.Left = nil
        
        if i+1 < len(nodeList) {
            n.Right = nodeList[i+1]
        } else {
            n.Right = nil 
        }
    }
    
    
    if len(nodeList) == 0 {
        return nil
    }
    return nodeList[0]
}

func helper(root *TreeNode, lst *[]*TreeNode) {
    if root == nil {
        return
    }
    
    helper(root.Left, lst)
    
    *lst = append(*lst, root)
    
    helper(root.Right, lst)
}