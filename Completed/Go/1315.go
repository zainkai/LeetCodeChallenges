// first
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumEvenGrandparent(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return countGrandChildren(root) + sumEvenGrandparent(root.Left) + sumEvenGrandparent(root.Right)
}

func countGrandChildren(grandParent *TreeNode) int {
	if grandParent.Val%2 == 1 {
		return 0
	}

	sumGrandChildren := 0
	for _, parent := range []*TreeNode{grandParent.Left, grandParent.Right} {
		if parent == nil {
			continue
		}
		if parent.Left != nil {
			sumGrandChildren += parent.Left.Val
		}
		if parent.Right != nil {
			sumGrandChildren += parent.Right.Val
		}
	}

	return sumGrandChildren
}

// second
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumEvenGrandparent(root *TreeNode) int {
	return countGrandChildren(root, nil, nil)
}

func countGrandChildren(child, parent, grandParent *TreeNode) int {
	if child == nil {
		return 0
	}

	sum := 0
	if grandParent != nil && grandParent.Val%2 == 0 {
		sum += child.Val
	}

	return sum + countGrandChildren(child.Left, child, parent) + countGrandChildren(child.Right, child, parent)
}