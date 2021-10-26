import (
    "strings"
    "strconv"
)

var (
    moves = [][]int{
        {1,0},
        {-1,0},
        {0,1},
        {0,-1},
    }
)

func numDistinctIslands(grid [][]int) int {
    res := 0
    islandSerial := map[string]bool{}
    
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[i]); j++ {
            if grid[i][j] == 0 {
                continue
            }
            sb := &strings.Builder{}
            helper(grid, i, j, 0,0, sb)
            
            if key := (*sb).String(); !islandSerial[key] {
                res++
                islandSerial[key] = true
            }
        }
    }
    return res
}

func helper(grid [][]int, x,y int, sX, sY int, sb *strings.Builder) {
    if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[x]) {
        return 
    } else if grid[x][y] == 0 {
        return
    }
    
    grid[x][y] = 0
    tmpX := strconv.Itoa(sX)
    tmpY := strconv.Itoa(sY)
    (*sb).WriteString(tmpX+","+tmpY+"|")
    for _, move := range moves {
        helper(grid, 
               x + move[0], 
               y + move[1],
               sX + move[0],
               sY + move[1],
               sb,
        )
    }
}
