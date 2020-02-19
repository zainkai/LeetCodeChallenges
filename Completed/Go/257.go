import (
	"strconv"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func binaryTreePaths(root *TreeNode) []string {
	results := make([]string, 0)
	if root == nil {
		return results
	}

	initalPath := strconv.Itoa(root.Val)
	if root.Left == nil && root.Right == nil {
		results = append(results, initalPath)
	} else {
		helper(root.Left, initalPath, &results)
		helper(root.Right, initalPath, &results)
	}

	return results
}

func helper(root *TreeNode, path string, rootToLeaf *[]string) {
	if root == nil {
		return
	}

	newPath := path + "->" + strconv.Itoa(root.Val)
	if root.Left == nil && root.Right == nil {
		*rootToLeaf = append(*rootToLeaf, newPath)
	}

	helper(root.Left, newPath, rootToLeaf)
	helper(root.Right, newPath, rootToLeaf)
}