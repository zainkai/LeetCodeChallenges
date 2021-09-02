import (
    "strings"
)

func findOcurrences(text string, first string, second string) []string {
    res := []string{}
    arr := strings.Split(text, " ")
    
    for i := range arr {
        if i+2 < len(arr) && arr[i] == first && arr[i+1] == second {
            third := arr[i+2]
            res = append(res, third)
        }
    }
    return res
}
