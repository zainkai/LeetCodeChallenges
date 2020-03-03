var numMap = map[byte][]string{
    '2': []string{"a","b","c"},
    '3': []string{"d","e","f"},
    '4': []string{"g","h","i"},
    '5': []string{"j","k","l"},
    '6': []string{"m","n","o"},
    '7': []string{"p","q","r","s"},
    '8': []string{"t","u","v"},
    '9': []string{"w","x","y","z"},
}

func letterCombinations(digits string) []string {
    result := []string{}
    if len(digits) == 0 {
        return result
    }
    
    helper(digits, 0, "", &result)
    
    return result
}

func helper(digits string, pos int, builder string, result *[]string) {
    if len(builder) == len(digits) {
        *result = append(*result, builder)
        return
    }
    
    b := digits[pos]
    for _, digit := range numMap[b]  {
        helper(digits, pos+1, builder+digit, result)
    }
}