const (
    COMP_LT = "LT"
    COMP_GT = "GT"
    MAX_TEAM_SIZE = 3
)

type frame struct {
    val int
    teamLen int
    comp string
}

func numTeams(rating []int) int {
    memo := map[frame]int{}
    
    return (helper(rating, frame{0, 0, COMP_GT}, memo)+
            helper(rating, frame{1<<63-1, 0, COMP_LT}, memo))
}

func helper(rating []int, f frame, memo map[frame]int) int {
    if res, ok := memo[f]; ok {
        return res
    }
    
    val, teamLen, comp := f.val, f.teamLen, f.comp
    if teamLen == MAX_TEAM_SIZE {
        return 1
    }
    
    result := 0
    for i, v := range rating {
        if comp == COMP_LT && v < val {
            result += helper(rating[i:], frame{v, teamLen+1, comp}, memo)   
        } else if comp == COMP_GT && v > val {
            result += helper(rating[i:], frame{v, teamLen+1, comp}, memo)  
        }
    }
    
    memo[f] = result
    return result
}
