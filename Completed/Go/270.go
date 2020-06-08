/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
import (
    "math"
)

type Data struct {
    diff float64
    num int
}

func closestValue(root *TreeNode, target float64) int {
    d := Data{math.MaxFloat64, 0}
    helper(root, target, &d)
    
    return d.num
}

func helper(root *TreeNode, target float64, d *Data) {
    if root == nil {
        return
    }
   
    f := float64(root.Val)
    if v := abs(f- target); v < (*d).diff {
        (*d).diff = v
        (*d).num = root.Val
    }
    
    helper(root.Left, target, d)
    helper(root.Right, target, d)
}

func abs(x float64) float64 {
    if x < 0 {
        return -1*x
    }
    return x
}

