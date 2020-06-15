type frame struct {
    row, col int
}

func minimumTotal(triangle [][]int) int {
    visited := map[frame]int{}
    
    return helper(frame{0,0}, triangle, visited)
}

func helper(f frame, triangle [][]int, visited map[frame]int) int {
    if res, ok := visited[f]; ok {
        return res
    } else if f.row >= len(triangle) {
        return 0
    } else if f.col < 0 || f.col >= len(triangle[f.row]) {
        return 1<<62-1
    }
    
    val := triangle[f.row][f.col]
    visited[f] = min(
        //helper(frame{f.row+1, f.col-1}, triangle, visited),
        helper(frame{f.row+1, f.col}, triangle, visited),
        helper(frame{f.row+1, f.col+1}, triangle, visited),
    ) + val
    
    return visited[f]
}


func min(el ...int) int {
    res := el[0]
    for _, val := range el {
        if val < res {
            res = val
        }
    }
    return res
}