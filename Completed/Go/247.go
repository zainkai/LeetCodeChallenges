var (
	strobMap = map[byte]byte{
		'1':'1',
		'6':'9',
		'8':'8',
        '9':'6',
        '0':'0',
    }
)


func findStrobogrammatic(n int) []string {
    output := []string{}
    
    
	
    if n == 1 {
        return []string{"0", "1", "8"}
    } else if n % 2 == 0 {
		helper("", n, &output)
    } else {
        helper("1", n-1, &output)
        helper("8", n-1, &output)
        helper("0", n-1, &output)
    }
	return output
}


func helper(str string, n int, output *[]string) {
	if n < 0 {
        return
    } else if n == 0 && str[0] != '0' { 
        *output = append(*output, str)
        return
    }

    for l, r := range strobMap {
        newStr := string(l) + str + string(r)
        helper(newStr, n-2, output)
    }
}


