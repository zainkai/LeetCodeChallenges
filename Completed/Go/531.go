func findLonelyPixel(picture [][]byte) int {
    rows := map[int]int{}
    cols := map[int]int{}
    
    for i := 0; i < len(picture); i++ {
        for j := 0; j < len(picture[i]); j++ {
            if picture[i][j] == 'B' {
                rows[i]++
                cols[j]++
            }
        }
    }
    
    res := 0
    for i := 0; i < len(picture); i++ {
        for j := 0; j < len(picture[i]); j++ {
            if picture[i][j] == 'B' && rows[i] == 1 && cols[j] == 1 { 
                res++
            }
        }
    }
    
    return res
}
