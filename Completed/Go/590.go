/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

 func postorder(root *Node) []int {
    stk := []*Node{root}
    res := []int{}
    
    
    for len(stk) > 0 {
        root = stk[len(stk)-1]
        stk = stk[:len(stk)-1]
        
        if root == nil {
            continue
        }
        
        res = append(res, root.Val)
        for i := 0; i < len(root.Children); i++ {
            child := root.Children[i]
            stk = append(stk, child)
        }
    }
    
    reverse(res)
    return res
}

func reverse(arr []int) {
    i := 0
    j := len(arr)-1
    for i < j {
        arr[i], arr[j] = arr[j], arr[i]
        i++
        j--
    }
}