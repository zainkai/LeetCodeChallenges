// back tracking instead of memo list
func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == word[0] && backtrack(board, word, 0, i, j) {
				return true
			}
		}
	}
	return false
}

func backtrack(board [][]byte, word string, pos int, i int, j int) bool {
	if i < 0 || i > len(board)-1 {
		return false
	} else if j < 0 || j > len(board[i])-1 {
		return false
	} else if pos == len(word)-1 && board[i][j] == word[pos] {
		return true
	} else if board[i][j] != word[pos] {
		return false
	}

	tmp := board[i][j]
	board[i][j] = byte(0)

	if backtrack(board, word, pos+1, i-1, j) || backtrack(board, word, pos+1, i+1, j) || backtrack(board, word, pos+1, i, j+1) || backtrack(board, word, pos+1, i, j-1) {
		board[i][j] = tmp
		return true
	}

	board[i][j] = tmp
	return false
}