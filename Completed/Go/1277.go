func countSquares(matrix [][]int) int {
	squares := 0
 	for i := len(matrix)-2; i >= 0; i-- {
		for j := len(matrix[i])-2; j >= 0; j-- {
			if matrix[i][j] == 0 {
				continue
            }

        helper(matrix, i, j)
        }
    }

    for _, arr := range matrix {
        for _, val := range arr {
            squares += val
        }
    }

    return squares
}

func helper(matrix [][]int, i, j int) {
	right := matrix[i][j+1]
	down := matrix[i+1][j]
	adj := matrix[i+1][j+1]

	matrix[i][j] = min(right, down, adj) +1
}

func min(nums... int) int {
    num := nums[0]
    for _, val := range nums {
        if val < num {
            num = val
        }
    }
    return num
}

