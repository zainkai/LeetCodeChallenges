var moves = [][]int{
    []int{0,1},
    []int{0,-1},
    []int{1,0},
    []int{-1,0},
}

func exist(board [][]byte, word string) bool {
    for i := range board {
        for j := range board[i] {
            if helper(i,j, board, word) {
                return true
            }
        }
    }
    return false
}

func helper(x,y int, board [][]byte, word string) bool {
    if len(word) == 0 {
        return true
    } else if x < 0 || y < 0 || x >= len(board) || y >= len(board[x]) {
        return false
    } else if top := word[0]; top != board[x][y] {
        return false
    }
    
    
    tmp := board[x][y]
    board[x][y] = byte(0)
    for _, move := range moves {
        if helper(x+move[0], y+move[1], board, word[1:]) {
            return true
        }
    }
    board[x][y] = tmp
    
    return false
}
