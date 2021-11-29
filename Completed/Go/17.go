var conv = map[byte][]byte{
    '2': []byte{'a','b','c'},
    '3': []byte{'d','e','f'},
    '4': []byte{'g','h','i'},
    '5': []byte{'j','k','l'},
    '6': []byte{'m','n','o'},
    '7': []byte{'p','q','r','s'},
    '8': []byte{'t','u','v'},
    '9': []byte{'w','x','y','z'},
}

func letterCombinations(digits string) []string {
    if len(digits) == 0 {
        return []string{}
    }
    
    res := []string{}
    helper(digits, []byte{}, &res)
    
    return res
}

func helper(digits string, path []byte, res *[]string) {
    if len(digits) == 0 {
        *res = append(*res, string(path))
        return
    }
    
    
    digit := digits[0]
    digits = digits[1:]
    for _, b := range conv[digit] {
        path = append(path, b)
        helper(digits, path, res)
        path = path[:len(path)-1]
    }
}
