type Point struct {
    i, j, lineType int
}

func longestLine(M [][]int) int {
    memo := map[Point]int{}
    
    maxLen := 0
    for i := 0; i < len(M); i++ {
        for j := 0; j < len(M[i]); j++ {
            for lineType := 0; lineType < 4; lineType++ {
                p := Point{i, j, lineType}
                if _, ok := memo[p]; !ok {
                    maxLen = max(maxLen, helper(M, p, memo))
                }
            }
            
        }
    }
    
    return maxLen
}

func helper(M [][]int, p Point, memo map[Point]int) int {
    if res, ok := memo[p]; ok {
        return res
    } else if p.i < 0 || p.j < 0 || p.i >= len(M) || p.j >= len(M[p.i]) {
        return 0
    } else if M[p.i][p.j] == 0 {
        return 0
    }
    
    if p.lineType == 0 { //horizontal
        return 1+ helper(M, Point{p.i+1, p.j, p.lineType}, memo)
    } else if p.lineType == 1 { // vertical
        return 1+ helper(M, Point{p.i, p.j-1, p.lineType}, memo)
    } else if p.lineType == 2 { // diag
        return 1+ helper(M, Point{p.i-1, p.j+1, p.lineType}, memo)
    } else {
        return 1+ helper(M, Point{p.i-1, p.j-1, p.lineType}, memo)
    }
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
