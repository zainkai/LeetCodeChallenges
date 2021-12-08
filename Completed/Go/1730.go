type point struct {
  x, y int
}

var moves = [][]int{
  {1,0},
  {-1,0},
  {0,1},
  {0,-1},
}

func getFood(grid [][]byte) int {
  start := point{0,0}
  
  findStart:
  for i := 0; i < len(grid); i++ {
    for j := 0; j < len(grid[i]); j++ {
      if grid[i][j] == '*' {
        start = point{i,j}
        break findStart
      }
    }
  }
  
  step := 0
  q := []point{start}
  for len(q) > 0 {
    length := len(q)
    
    for i := 0; i < length; i++ {
      top := q[0]
      q = q[1:]
      if !inBounds(grid, top) || grid[top.x][top.y] == 'X' {
        continue
      } else if grid[top.x][top.y] == '#' {
        return step
      }
      
      grid[top.x][top.y] = 'X'
      for _, mov := range moves {
        q = append(q, point{
          top.x+mov[0],
          top.y+mov[1],
        })
      }
    }
    step++
  }
  
  return -1
}

func inBounds(grid [][]byte, p point) bool {
  return !(p.x < 0 || p.y < 0 || p.x >= len(grid) || p.y >= len(grid[p.x]))
}
