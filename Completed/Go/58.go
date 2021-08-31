import (
    "strings"
)

func lengthOfLastWord(s string) int {
    words := strings.Split(s, " ")
    
    fmt.Println(words)
    for i := len(words)-1; i >= 0; i-- {
        if len(words[i]) == 0 {
            continue
        } else {
            return len(words[i])
        }
    }
    return 0
}
