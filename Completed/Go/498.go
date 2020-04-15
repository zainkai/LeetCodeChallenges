func findDiagonalOrder(matrix [][]int) []int {
    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return []int{}
    }
    
    diagonals := [][]int{}
    
    for i := 0; i < len(matrix[0]); i++ {
        diagonals = append(diagonals, getDiagonal(matrix, 0, i))
    }
    
    for i := 1; i < len(matrix); i++ {
        diagonals = append(diagonals, getDiagonal(matrix, i, len(matrix[i])-1))
    }
    
    res := []int{}
    for i, arr := range diagonals {
        if i % 2 == 0 {
            reverseInts(arr)
        }
        res = append(res, arr...)
    }
    return res
}

func getDiagonal(matrix [][]int, x, y int) []int {
    res := []int{}
    for y >= 0 && x < len(matrix) {
        res = append(res, matrix[x][y])
        x++
        y--
    }
    return res
}

func reverseInts(a []int) {
    i, j := 0, len(a)-1
    for i < j {
        a[i], a[j] = a[j], a[i]
        
        i++
        j--
    }
}