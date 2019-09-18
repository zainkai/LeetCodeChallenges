/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type level struct {
	sum   float64
	nodes float64
}
type levelMap map[int]*level

func averageOfLevels(root *TreeNode) []float64 {
	m := levelMap{}
	maxDepth := getLevels(0, root, &m)

	result := make([]float64, 0)
	for i := 0; i < maxDepth; i++ {
		val := m[i].sum / m[i].nodes
		result = append(result, val)
	}

	return result
}

func getLevels(depth int, root *TreeNode, m *levelMap) int {
	if root == nil {
		return depth
	}

	if _, ok := (*m)[depth]; !ok {
		(*m)[depth] = new(level)
	}

	(*m)[depth].sum += float64(root.Val)
	(*m)[depth].nodes += 1

	a := getLevels(depth+1, root.Left, m)
	b := getLevels(depth+1, root.Right, m)

	if a > b {
		return a
	}
	return b
}