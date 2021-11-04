func minOperations(logs []string) int {
    stk := []string{}
    for _, dir := range logs {
        if dir == "./" {
            continue
        } else if dir == "../" {
            if len(stk) > 0 {
                stk = stk[:len(stk)-1]
            }
        } else {
            stk = append(stk, dir)
        }
    }
    
    return len(stk)
}
