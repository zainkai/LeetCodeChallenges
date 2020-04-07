/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func flatten(root *TreeNode) {
	nArr := []*TreeNode{}
	preOrder(root, &nArr)

	for i := 0; i < len(nArr)-1; i++ {
		nArr[i].Right = nArr[i+1]
	}
}

func preOrder(root *TreeNode, nodeArr *[]*TreeNode) {
	if root == nil {
		return
	}

	*nodeArr = append(*nodeArr, root)

	preOrder(root.Left, nodeArr)
	preOrder(root.Right, nodeArr)

	root.Left = nil
	root.Right = nil
}