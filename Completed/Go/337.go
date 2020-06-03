/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 type frame struct {
    root *TreeNode
    rob bool
}

func rob(root *TreeNode) int {
    memo := map[frame]int{}
    
    return max(
        helper(root, false, memo),
        helper(root, true, memo),
    )
}

func helper(root *TreeNode, doRob bool, memo map[frame]int) int {
    f := frame{root, doRob}
    if res, ok := memo[f]; ok {
        return res
    } else if root == nil {
        return 0
    }
    
    value := 0
    if doRob {
        value += root.Val
    }
    
    memo[f] = max(
        value + helper(root.Left, !doRob, memo) + helper(root.Right, !doRob, memo),
        helper(root.Left, doRob, memo) + helper(root.Right, doRob, memo),
    )
    
    return memo[f]
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}