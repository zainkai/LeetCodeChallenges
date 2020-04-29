import (
    "strings"
)

type EAddr struct {
    local, domain string
}

func numUniqueEmails(emails []string) int {
    uniqueEmails := map[EAddr]bool{}
    
    for _, email := range emails {
        uniqueEmails[initEAddr(email)] = true
    }
    
    return len(uniqueEmails)
}

func initEAddr(email string) EAddr {
    ea := EAddr{}
    
    tok := strings.Split(email, "@")

    ea.domain = tok[1]
    
    local := tok[0]
    sb := strings.Builder{}
    for _, r := range local {
        if r == '.' {
            continue
        } else if r == '+' {
            break
        } else {
            sb.WriteRune(r)
        }
    }
    ea.local = sb.String()
    
    return ea
}