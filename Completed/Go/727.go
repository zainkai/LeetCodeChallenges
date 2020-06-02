type Point struct {
    i, j int
}

func minWindow(S string, T string) string {
    l, r := 0, 1<<63-1
    
    memo := map[Point]int{}
    for i := 0; i < len(S); i++ {
        rTmp := helper(S, T, Point{i, 0}, memo)
        if rTmp < r-l {
            l = i
            r = i+rTmp
        }
    }
    
    if r >= 0 && r <= len(S) {
        return S[l:r]
    }
    return ""
}

func helper(S, T string, p Point, memo map[Point]int) int {
    if val, ok := memo[p]; ok {
        return val
    } else if p.j == len(T) {
        return 0
    } else if p.i >= len(S) {
        return 1 << 62-1
    }
    
    if S[p.i] == T[p.j] {
        memo[p] = 1+helper(S,T, Point{p.i+1, p.j+1}, memo)
    } else {
        memo[p] = 1+helper(S,T, Point{p.i+1, p.j}, memo)
    }
    return memo[p]
}