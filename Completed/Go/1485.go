/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Random *Node
 * }
 */

func copyRandomBinaryTree(root *Node) *NodeCopy {
    memo := map[*Node]*NodeCopy{}
    
    return helper(root, memo)
}

func getOrCreateNode(root *Node, memo map[*Node]*NodeCopy) *NodeCopy {
    if res, ok := memo[root]; ok {
        return res
    }
    
    memo[root] = &NodeCopy{
        root.Val,
        nil,
        nil,
        nil,
    }
    
    return memo[root]
}

func helper(root *Node, memo map[*Node]*NodeCopy) *NodeCopy {
    if root == nil {
        return nil
    }
    
    tRoot := getOrCreateNode(root, memo)
    tRoot.Left = helper(root.Left, memo)
    tRoot.Right = helper(root.Right, memo)
    if root.Random != nil {
        tRoot.Random = getOrCreateNode(root.Random, memo)
    }
    
    return tRoot
}
