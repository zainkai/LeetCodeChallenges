type frame struct {
    i, j int
}

func minDistance(word1 string, word2 string) int {
    memo := map[frame]int{}
    return helper(word1, word2, frame{ 0, 0 }, memo)
}

func helper(w1, w2 string, f frame, memo map[frame]int) int {
    if res, ok := memo[f]; ok {
        return res
    }
    
    i, j := f.i, f.j
    if i == len(w1) && j == len(w2) {
        return 0
    } else if i > len(w1) || j > len(w2) {
        return 1<< 63-1
    }
    
    if i < len(w1) && j< len(w2) && w1[i] == w2[j] {
        memo[f] = helper(w1, w2, frame{i+1, j+1}, memo)
        return memo[f]
    }
    
    memo[f] = 1 + min(
        helper(w1, w2, frame{i, j+1}, memo), // insert
        helper(w1, w2, frame{i+1, j}, memo), // delete
        helper(w1, w2, frame{i+1, j+1}, memo), // replace
    )
    
    return memo[f]
}

func min(el ...int) int {
    res := el[0]
    for _, v := range el {
        if v < res {
            res = v
        }
    }
    return res
}




