/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type frame struct {
    node *TreeNode
    r int
}

func verticalOrder(root *TreeNode) [][]int {
    res := [][]int{}
    if root == nil {
        return res
    }
    
    q := []frame{
        {root, 0},
    }
    leftMostVert := 0
    rightMostVert := 0
    
    vertMap := map[int][]int{}
    for len(q) > 0 {
        node, r := q[0].node, q[0].r
        q = q[1:]
        
        leftMostVert = min(leftMostVert, r)
        rightMostVert = max(rightMostVert, r)
        
        vertMap[r] = append(vertMap[r], node.Val)
        if node.Left != nil {
            q = append(q, frame{node.Left, r-1})
        }
        if node.Right != nil {
            q = append(q, frame{node.Right, r+1})
        }
    }
    
    for i:=leftMostVert; i <= rightMostVert; i++ {
        res = append(res, vertMap[i])
    }
    return res
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func max(a,b int) int {
    if a > b {
        return a
    }
    return b
}
