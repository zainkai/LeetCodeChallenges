type Point struct {
	a, b int
}

func maxUncrossedLines(A, B []int) int {
	memo := map[Point]int{}

    return helper(A,B, Point{-1, -1}, memo)
}

func helper(A, B []int, pt Point, memo map[Point]int) int {
	if val, ok := memo[pt]; ok {
		return val
    }

    for i := pt.a+1; i < len(A); i++ {
        for j := pt.b+1; j < len(B); j++ {
            if A[i] == B[j] {
                memo[pt] = max(memo[pt], 1)
                memo[pt] = max(memo[pt], helper(A,B, Point{i,j}, memo)+1)
            }
        }
    }
    return memo[pt]
}

func max(a, b int) int {
	if a > b {
		return a
    }
    return b
}
