var moves = [][]int{
    {0,1},
    {0,-1},
    {1,0},
    {-1, 0},
}

func solve(board [][]byte)  {
    for i := range board {
        helper(i, 0, board, 'O', 'R')
        helper(i, len(board[i])-1, board, 'O', 'R')
    }
    for i := range board[0] {
        helper(0, i, board, 'O', 'R')
        helper(len(board)-1, i, board, 'O', 'R')
    }
    
    for i := range board {
        for j := range board[i] {
            helper(i,j, board, 'O', 'X')
        }
    }
    
    for i := range board {
        for j := range board[i] {
            helper(i,j, board, 'R', 'O')
        }
    }
}

func helper(x, y int, board [][]byte, target byte, replacement byte) {
    if x < 0 || y < 0 || x >= len(board) || y >= len(board[x]) {
        return
    } else if board[x][y] != target {
        return
    }
    board[x][y] = replacement
    
    for _, mov := range moves {
        helper(x+mov[0], y+mov[1], board, target, replacement)
    }
}
