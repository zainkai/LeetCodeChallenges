func longestIncreasingPath(matrix [][]int) int {
    dp := make([][]int, len(matrix))
    for i := range matrix {
        dp[i] = make([]int, len(matrix[i]))
    }
    
    longest := 0
    for i := 0; i < len(matrix); i++ {
        for j := 0; j < len(matrix[i]); j++ {
            longest = max(dfs(matrix, dp, i, j, -1<<63), longest)
        }
    }
    
    return longest
}

var moves = [][]int{{-1,0},{1,0},{0,-1},{0,1}}

func dfs(matrix, dp [][]int, i, j, parent int) int {
    if i < 0 || j < 0 || i >= len(dp) || j >= len(dp[i]) {
        return 0
    } else if parent >= matrix[i][j] {
        return 0
    } else if dp[i][j] > 0 {
        return dp[i][j]
    }
    
    dp[i][j] =1
    
    for _, mov := range moves {
        r, c := i+mov[0], j+mov[1]
        
        dp[i][j] = max(
            dp[i][j],
            1+dfs(matrix, dp, r, c, matrix[i][j]),
        )    
    }
    
    return dp[i][j]
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}