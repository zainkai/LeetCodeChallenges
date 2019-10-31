/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	newNode := &TreeNode{val, nil, nil}
	if root == nil {
		return newNode
	}

	temp := root
	for {
		if val < temp.Val {
			if temp.Left == nil {
				temp.Left = newNode
				return root
			} else {
				temp = temp.Left
			}
		} else {
			if temp.Right == nil {
				temp.Right = newNode
				return root
			} else {
				temp = temp.Right
			}
		}
	}

	return root
}