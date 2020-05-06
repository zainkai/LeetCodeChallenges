import (
    "strconv"
)

func numDecodings(s string) int {
    if len(s) == 0 {
        return 0
    }
    
	memo := map[int]int{
		len(s): 1,
	}

	return helper(s, 0, memo)
}

func helper(s string, pos int, memo map[int]int) int {
	if val, ok := memo[pos]; ok {
		return val
	} else if len(s) < pos {
		return 0
    } else if s[pos] == '0' {
        return 0
    }

	if pos < len(s) && isValid(string(s[pos])) {
		memo[pos] += helper(s, pos+1, memo)
	}

	if pos+1 < len(s) && isValid(string(s[pos:pos+2])) {
		memo[pos] += helper(s, pos+2, memo)
	}

	return memo[pos]
}

func isValid(s string) bool {
    fmt.Println(s)
    v, _ := strconv.Atoi(s)
    return v >= 1 && v <= 26
}
