type data struct {
	parent int
	depth  int
	target int
}

func isCousins(root *TreeNode, x, y int) bool {
	if x == y || root == nil {
		return false
	}

	xData := data{
		-1, 0, x,
	}
	yData := data{
		-1, 0, y,
	}

	helper(root, nil, &xData, &yData, 0)

    return xData.depth == yData.depth && xData.parent != yData.parent
}

func helper(root, parent *TreeNode, x, y *data, depth int) {
    if root == nil {
        return
    }
    
    if root.Val == x.target {
        x.depth = depth
        if parent != nil {
            x.parent = parent.Val
        }
    }
    if root.Val == y.target {
        y.depth = depth
        if parent != nil {
            y.parent = parent.Val
        }
    }
    
    helper(root.Left, root, x, y, depth+1)
    helper(root.Right, root, x, y, depth+1)
}