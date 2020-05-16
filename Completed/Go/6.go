import (
    "strings"
)

func convert(s string, numRows int) string {
    if numRows == 1 {
        return s
    }
    
    rArr := make([][]rune, numRows)
    
    row := 0
    isUp := false
    for _, r := range s { 
        rArr[row] = append(rArr[row], r)
        
        if row == 0 || row == numRows -1 {
            isUp = !isUp
        }
        
        if isUp {
            row++
        } else {
            row--
        }
    }
    
    sb := strings.Builder{}
    for _, arr := range rArr {
        for _, r := range arr {
            sb.WriteRune(r)
        }
    }
    
    return sb.String()
}