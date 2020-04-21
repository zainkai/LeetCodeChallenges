func diagonalSort(mat [][]int) [][]int {
    if len(mat) == 0 || len(mat[0]) == 0 {
        return mat
    }
    
    for i := 0; i < len(mat[0]); i++ {
        quickSort(mat, 0,i, len(mat), len(mat[0]))
    }
    for i := 0; i < len(mat); i++ {
        quickSort(mat, i,0, len(mat), len(mat[i]))
    }
    return mat
}

/*func quickSort(matrix [][]int, xlow, ylow, xhigh, yhigh int) {
    if xlow >= xhigh || ylow >= yhigh {
        return
    }
    
    pivotEle := matrix[xhigh][yhigh]
    xtmp, ytmp := xlow, ylow
    
    for x, y := xlow, ylow; x < xhigh && y < yhigh; x, y = x+1, y+1 {
        if matrix[x][y] < pivotEle {
            matrix[x][y], matrix[xtmp][ytmp] = matrix[xtmp][ytmp], matrix[x][y]
            
            xtmp++
            ytmp++
        }
    }
    
    matrix[xhigh][yhigh], matrix[xtmp][ytmp] = matrix[xtmp][ytmp], matrix[xhigh][yhigh]
    
    quickSort(matrix, xlow, ylow, xtmp-1, ytmp-1)
    quickSort(matrix, xtmp+1, ytmp+1, xhigh, yhigh)
}*/

func quickSort(matrix [][]int, xlow, ylow, xhigh, yhigh int) {
    arr := []int{}
    for x, y := xlow, ylow; x < xhigh && y < yhigh; x, y = x+1, y+1 {
        arr = append(arr , matrix[x][y])
    }
    sort.Ints(arr)
    i:= 0
    for x, y := xlow, ylow; x < xhigh && y < yhigh; x, y = x+1, y+1 {
        matrix[x][y] = arr[i]
        i++
    }
}