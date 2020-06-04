const (
    bucketSize = 9
)

func isValidSudoku(board [][]byte) bool {
    rows := genBuckets()
    cols := genBuckets()
    squares := genBuckets()
    
    for i := 0; i < len(board); i++ {
        for j := 0; j < len(board[i]); j++ {
            val := int(board[i][j] - '0') -1
            if val < 0 || val > 9 {
                continue
            }
            
            sqIdx := (i / 3 ) * 3 + j / 3
            // visit for each bucket
            if !rows[i][val] && !cols[j][val] && !squares[sqIdx][val] {
                rows[i][val] = true 
                cols[j][val] = true
                squares[sqIdx][val] = true
            } else {
                return false
            }
        }
    }
    
    return true
}

func genBuckets() map[int][]bool{
    bucket := make(map[int][]bool)
    for i := 0; i < bucketSize; i++ {
        bucket[i] = make([]bool, bucketSize)
    }
    
    return bucket
}
