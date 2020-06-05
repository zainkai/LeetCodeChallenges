func flipAndInvertImage(A [][]int) [][]int {
    for row := range A {
        A[row] = flipAndInvert(A[row])
    }
    return A
}

func flipAndInvert(row []int) []int {
    i, j := 0, len(row)-1
    
    for i <= j {
        row[i], row[j] = flip(row[j]), flip(row[i])
        
        
        i++
        j--
    }
    
    fmt.Println(row)
    return row
}

func flip(i int) int {
    if i == 0 {
        return 1
    }
    return 0
}
