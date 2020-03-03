func generateParenthesis(n int) []string {
    result := []string{}
    dfs(&result, n, "", 0, 0)
    
    return result
}

func dfs(result *[]string, n int, s string, open int, closing int) {
    if open == n && closing == n {
        *result = append(*result, s)
    } else if open > n || closing > n {
        return
    }
    
    if open < n {
        dfs(result, n, s+"(", open+1, closing)
    }
    
    if closing < open {
        dfs(result, n, s+")", open, closing+1)
    }
}