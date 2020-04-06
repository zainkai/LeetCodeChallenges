/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
    del := map[int]bool{}
	for _, v := range to_delete {
		del[v] = true
    }

    forest := make([]*TreeNode, 0, 10)
    
    if root != nil && !del[root.Val] {
        forest = append(forest, root)
    }
    
    helper(root, del, &forest)
    return forest
}

func helper(root *TreeNode, delMap map[int]bool, forest *[]*TreeNode) *TreeNode {
    if root == nil {
        return nil
    }

    if _, found := delMap[root.Val]; found {
        if root.Left != nil && !delMap[root.Left.Val] {
            *forest = append(*forest, root.Left)   
        }
        if root.Right != nil && !delMap[root.Right.Val] {
            *forest = append(*forest, root.Right)
        }
        
        helper(root.Left, delMap, forest)
        helper(root.Right, delMap, forest)
        return nil
    }

    root.Left = helper(root.Left, delMap, forest)
    root.Right = helper(root.Right, delMap, forest)
    return root
}




