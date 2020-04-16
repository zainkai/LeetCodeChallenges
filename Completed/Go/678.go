type memorize struct {
    s string
    brackets int
}

func checkValidString(s string) bool {
    memo := map[memorize]bool{}
    helper(s, 0, memo)
    
    for _, res := range memo {
        if res {
            return true
        }
    }
    return false
}

func helper(s string, brackets int, memo map[memorize]bool) {
    if _, ok := memo[memorize{s, brackets}]; ok {
        return
    }
    
    for i:= 0 ;i < len(s); i++ {
        if brackets < 0 {
            memo[memorize{s, brackets}] = false
            return
        }
        
        if s[i] == '(' {
            brackets++
        } else if s[i] == ')' {
            brackets--
        } else {
            helper(s[i+1:], brackets-1, memo)
            helper(s[i+1:], brackets+1, memo) 
        }
    }
    memo[memorize{s, brackets}] = brackets == 0
}