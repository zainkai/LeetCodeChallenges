var board = [][]int{
    {1,1,1},
    {1,1,1},
    {1,1,1},
    {0,1,0},
}

type frame struct {
    x, y, n int
}

func knightDialer(N int) int {
    visited := map[frame]int{}
    
    uniquePaths := 0
    for i := 0; i < len(board); i++ {
        for j := 0; j < len(board[i]); j++ {
            uniquePaths += dfs(frame{i, j, N}, visited)
        }
    }
    
    return uniquePaths % 1000000007
}

func dfs(f frame, visited map[frame]int) int {
    if paths, ok := visited[f]; ok {
        return paths
    } else if f.x < 0 || f.y < 0 || f.x >= len(board) || f.y >= len(board[f.x]) {
        return 0
    } else if board[f.x][f.y] == 0 {
      return 0  
    } else if f.n == 1 {
        return 1
    }
    
    moves := [][]int{{-2,-1},{-1,-2},{-2,1},{-1,2},{1,-2},{2,-1},{1,2},{2,1}}
    for _, move := range moves {
        visited[f] += dfs(frame{
            x: f.x +move[0],
            y: f.y +move[1],
            n: f.n -1,
        }, visited)  % 1000000007
    }
    
    return visited[f]
}

