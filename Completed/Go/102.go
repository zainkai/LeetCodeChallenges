/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type frame struct {
    height int
    n *TreeNode
}


func levelOrder(root *TreeNode) [][]int {
    res := [][]int{}
    if root == nil {
        return res
    }
    
    q := []frame{}
    Enque(&q, frame{
        height:0,
        n:root,
    })

    for len(q) > 0 {
        f := Deque(&q)
        if f.height >= len(res) {
            res = append(res, []int{})
        }
        
        res[f.height] = append(res[f.height], f.n.Val)
        
        if f.n.Right != nil {
            Enque(&q, frame{
                height: f.height+1,
                n: f.n.Right,
            })
        }
        if f.n.Left != nil {
            Enque(&q, frame{
                height: f.height+1,
                n: f.n.Left,
            })
        }
        
    }
    return res
}


// helpers
func Enque(q *[]frame, f frame) {
    *q = append(*q, f)
}

func Deque(q *[]frame) frame {
    top := (*q)[len(*q)-1]
    *q = (*q)[:len(*q)-1]
    
    return top
}

