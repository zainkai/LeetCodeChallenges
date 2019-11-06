import (
    "strings"
)


const Token string = "[.]"

func defangIPaddr(address string) string {
    strFactory := strings.Builder{}
    
    for _, r := range address {
        if r == '.' {
            strFactory.WriteString(Token)
        } else {
            strFactory.WriteRune(r)
        }
    }
    
    return strFactory.String()
}