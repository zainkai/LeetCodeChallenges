import (
	"strings"
)

func simplifyPath(path string) string {
	strArr := strings.Split(path, "/")
	
	stk := []string{}
	for _, s := range strArr {
		if strings.EqualFold(s, "") {
			continue
        } else if strings.EqualFold(s, ".") {
            continue
        } else if strings.EqualFold(s, "..") {
            if len(stk) > 0 {
                stk = stk[:len(stk)-1] // (POP) sub-slicing is O(1)
            }
        } else {
            stk = append(stk, s) // (PUSH) O(1) amortized
        }
    }

    sb := strings.Builder{}
    for _, p := range stk {
        sb.WriteString("/"+p)
    }
    
    if len(stk) == 0 {
        return "/"
    }
	return sb.String()
}

