var (
    strobs = map[string]string{
        "0":"0",
        "1":"1",
        "6":"9",
        "8":"8",
        "9":"6",
    }
    
    seeds = []string{"0", "1", "8"}
)

func findStrobogrammatic(n int) []string {
    if n == 1 {
        return seeds
    }
    
    res := []string{}
    if n % 2 == 0 {
        helper(n, "", &res)
    } else {
        for _, seed := range seeds {
            helper(n, seed, &res)
        }
    }
    
    return res
}

func helper(n int, s string, res *[]string) {
    if len(s) == n && s[0] != '0' {
        *res = append(*res, s)
        return
    } else if len(s) > n {
        return
    }
    
    for l, r := range strobs {
        helper(n, l+s+r, res)
    }
}
