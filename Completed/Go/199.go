/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
import (
	"container/list"
)

type node struct {
	level int
	ele   *TreeNode
}

func rightSideView(root *TreeNode) []int {
	q, maxDepth := list.New(), 0
	if root != nil {
		q.PushBack(node{0, root})
	}

	m := map[int]int{}

	for q.Len() > 0 {
		n := q.Front().Value.(node)
		maxDepth = max(maxDepth, n.level)

		m[n.level] = n.ele.Val

		if n.ele.Left != nil {
			q.PushBack(node{n.level + 1, n.ele.Left})
		}
		if n.ele.Right != nil {
			q.PushBack(node{n.level + 1, n.ele.Right})
		}
		q.Remove(q.Front())
	}

	if len(m) == 0 {
		return []int{}
	}

	res := make([]int, maxDepth+1)
	for lvl, val := range m {
		res[lvl] = val
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}