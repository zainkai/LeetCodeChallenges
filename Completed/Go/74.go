func searchMatrix(matrix [][]int, target int) bool {
    row, col := len(matrix)-1, 0
    for row >= 0 && col < len(matrix[row]) {
        if matrix[row][col] == target {
            return true
        }
        
        top := -1<<63
        if row-1 >= 0 {
            top = matrix[row-1][col]
        }
        
        if target <= top {
            row--
        } else {
            col++
        }
    }
    return false
}
