var isDigit = map[byte]int {
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

func calculate(s string) int {
    arr := []byte{}
    
    // trim spaces and traverse in in reverse
    for i := len(s)-1; i >= 0; i-- {
        if s[i] == byte(0) {
            continue
        } else {
            arr = append(arr, s[i])
        }
    }
    
    return helper(&arr)
}

func helper(arr *[]byte) int {
    stk := []int{}
    sign := byte('+')
    num := 0
    
    for len(*arr) > 0 {
        top := (*arr)[len(*arr)-1] // top
        *arr = (*arr)[:len(*arr)-1] // pop
        
        if v, ok := isDigit[top]; ok {
            num *= 10
            num += v
        } else if top == '(' {
            num = helper(arr)
        } else if top == '+' || top == '-' || top == ')'{
            if sign == '+' {
                stk = append(stk, num)
            } else if sign == '-' {
                stk = append(stk, -num)
            }
            
            sign = top
            num = 0
            if sign == ')' { // end recursive frame
                break
            }
        }
    }
    
    // add last num
    if sign == '+' {
        stk = append(stk, num)
    } else if sign == '-' {
        stk = append(stk, -num)
    }
    
    return arrSum(stk)
}

func arrSum(arr []int) int {
    res := 0
    for _, v := range arr {
        res += v
    }
    
    return res
}


