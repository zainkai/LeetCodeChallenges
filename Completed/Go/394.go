import (
    "strings"
)

var digitMap = map[byte]int{
    '0':0,
    '1':1,
    '2':2,
    '3':3,
    '4':4,
    '5':5,
    '6':6,
    '7':7,
    '8':8,
    '9':9,
}

func decodeString(s string) string {
    arr := []byte(s)
    return helper(&arr)
}

func helper(arr *[]byte) string {
    sb := strings.Builder{}
    mult := 0
    
    for len(*arr) > 0 {
        top := (*arr)[0]
        *arr = (*arr)[1:]
        
        if val, ok := digitMap[top]; ok {
            mult *= 10
            mult += val
        } else if top == '[' {
            str := helper(arr)
            sb.WriteString(strings.Repeat(str, mult))
            mult = 0
        } else if top == ']' {
            break
        } else {
            sb.WriteByte(top)
        }
    }
    
    return sb.String()
}
