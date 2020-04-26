func longestCommonSubsequence(t1 string, t2 string) int {
    memo := make(map[frame]int)
    
    return helper(t1,t2, frame{len(t1), len(t2)}, memo)
}

type frame struct {
    l1, l2 int
}

func helper(t1, t2 string, f frame, memo map[frame]int) int {
    if val, ok := memo[f]; ok {
        return val
    }
    
    if f.l1 == 0 || f.l2 == 0 {
        return memo[f]
    } else if t1[f.l1-1] == t2[f.l2-1] {
        x := frame{f.l1-1, f.l2-1}
        memo[f] = 1+helper(t1,t2, x, memo)
    } else {
        a := frame{f.l1-1, f.l2}
        b := frame{f.l1, f.l2-1}
        
        memo[f] = max(
            helper(t1,t2, a, memo),
            helper(t1,t2, b, memo),
        )
    }
    
    return memo[f]
}

func max(a,b int) int {
    if a > b {
        return a
    }
    return b
}