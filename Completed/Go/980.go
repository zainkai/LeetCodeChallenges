var moves = [][]int{
    {0,1},
    {0,-1},
    {1,0},
    {-1,0},
}

func uniquePathsIII(grid [][]int) int {
    start := []int{0,0}
    end := []int{0,0}
    empty := 1 // account for start point
    
    for i := range grid {
        for j := range grid[i] {
            if grid[i][j] == 0 {
                empty++
            } else if grid[i][j] == 1 {
                start = []int{i,j}
            } else if grid[i][j] == 2 {
                end = []int{i,j}
            }
        }
    }
    
    res := 0
    helper(grid, empty, start, end, &res)
    
    return res
}

func helper(grid [][]int, empty int, pos, end []int, res *int) {
    if pos[0] < 0 || pos[1] < 0 || pos[0] >= len(grid) || pos[1] >= len(grid[pos[0]]) {
        return
    } else if grid[pos[0]][pos[1]] == -1 {
        return
    } else if pos[0] == end[0] && pos[1] == end[1] {
        if empty == 0 {
            (*res)++
        }
        return
    }
    
    tmp := grid[pos[0]][pos[1]]
    grid[pos[0]][pos[1]] = -1
    
    for _, mov := range moves {
        helper(
            grid, 
            empty-1, 
            []int{ pos[0]+mov[0], pos[1]+mov[1] },
            end,
            res,
        )
    }
    
    grid[pos[0]][pos[1]] = tmp
} 
