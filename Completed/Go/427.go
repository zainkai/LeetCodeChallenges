/**
 * Definition for a QuadTree node.
 * type Node struct {
 *     Val bool
 *     IsLeaf bool
 *     TopLeft *Node
 *     TopRight *Node
 *     BottomLeft *Node
 *     BottomRight *Node
 * }
 */

func construct(grid [][]int) *Node {
    return helper(grid, 0, len(grid), 0, len(grid))
}

func helper(grid [][]int, rLo, rHi, cLo, cHi int) *Node {
    if isSame(grid, rLo, rHi, cLo, cHi) {
        val := grid[rLo][cLo] == 1
        return &Node{
            Val: val,
            IsLeaf: true,
            TopLeft: nil,
            TopRight: nil,
            BottomLeft: nil,
            BottomRight: nil,
        }
    } else {
        rMid := (rLo+rHi)/2
        cMid := (cLo+cHi)/2
        
        return &Node{
            Val: false,
            IsLeaf: false,
            TopLeft: helper(grid, rLo, rMid, cLo, cMid),
            TopRight: helper(grid, rLo, rMid, cMid, cHi),
            BottomLeft: helper(grid, rMid, rHi, cLo, cMid),
            BottomRight: helper(grid, rMid, rHi, cMid, cHi),
        }
    }
}

func isSame(grid [][]int, rLo, rHi, cLo, cHi int) bool {
    target := grid[rLo][cLo]
    
    for i := rLo; i < rHi; i++ {
        for j := cLo; j < cHi; j++ {
            if grid[i][j] != target {
                return false
            }
        }
    }
    
    return true
}
