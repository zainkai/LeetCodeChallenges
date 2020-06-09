/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxAncestorDiff(root *TreeNode) int {
    diff := -1<<63
    helper(root, &diff)

    return diff
}

func helper(root *TreeNode, diff *int) (int, int) {
    resMin := root.Val
    resMax := root.Val
    
    if root.Left != nil {
        tmin, tmax := helper(root.Left, diff)
        *diff = max(*diff, abs(root.Val-tmin))
        *diff = max(*diff, abs(root.Val-tmax))
        
        resMin = min(resMin, tmin)
        resMax = max(resMax, tmax)
    }
    if root.Right != nil {
        tmin, tmax := helper(root.Right, diff)
        *diff = max(*diff, abs(root.Val-tmin))
        *diff = max(*diff, abs(root.Val-tmax))
        
        resMin = min(resMin, tmin)
        resMax = max(resMax, tmax)
    }
        
    return resMin, resMax
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i <= j {
		return i
	}
	return j
}

func abs(x int) int {
    if x < 0 {
        return -1 *x
    }
    return x
}
