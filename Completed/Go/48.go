func rotate(matrix [][]int) {
	for row := range matrix {
		for col := 0; col < row; col++ {
			matrix[row][col], matrix[col][row] = matrix[col][row], matrix[row][col]
		}
	}

	for row := range matrix {
		reverseRow(matrix, row)
	}
}

func reverseRow(matrix [][]int, row int) {
	i, j := 0, len(matrix)-1
	for i < j {
		matrix[row][i], matrix[row][j] = matrix[row][j], matrix[row][i]
		i++
		j--
	}
}