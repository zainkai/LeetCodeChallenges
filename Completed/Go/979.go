/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func distributeCoins(root *TreeNode) int {
    moves := 0
    helper(root, &moves)
    
    return moves
}

func helper(root *TreeNode, moves *int) int {
    if root == nil {
        return 0
    }
    
    leftCoins := helper(root.Left, moves)
    *moves += abs(leftCoins)
    
    rightCoins := helper(root.Right, moves)
    *moves += abs(rightCoins)
    
    if root.Val == 1 {
        return leftCoins + rightCoins
    } else if root.Val == 0 {
        return -1 + leftCoins + rightCoins
    } else {
        return root.Val - 1 + leftCoins + rightCoins
    }
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}