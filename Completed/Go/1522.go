/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

import (
    "sort"
)

func diameter(root *Node) int {
    res := 0
    helper(root, &res)
    
    return res
}

func helper(root *Node, diam *int) int {
    if root == nil {
        return 0
    }
    
    res := make([]int, len(root.Children))
    maxSinglePath := 0
    for i, child := range root.Children {
        res[i] = helper(child, diam)
        maxSinglePath = max(res[i], maxSinglePath)
    }
    
    sort.Slice(res, func(i, j int) bool {
        return res[i] > res[j]
    })
    
    if len(res) >= 2 {
        *diam = max(*diam, res[0]+res[1])
    } else if len(res) == 1 {
        *diam = max(*diam, res[0])
    }
    
    return maxSinglePath+1
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
