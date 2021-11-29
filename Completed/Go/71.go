import "strings"

func simplifyPath(path string) string {
    stk := []string{}
    
    pathArr := strings.Split(path, "/")
    for _, elm := range pathArr {
        if elm == "" || elm == "." {
            continue
        } else if elm == ".." {
            if len(stk) > 0 {
                stk = stk[:len(stk)-1]
            }
        } else {
            stk = append(stk, elm)
        }
    }
    
    return "/" + strings.Join(stk, "/")
}
