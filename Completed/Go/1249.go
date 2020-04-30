import (
    "strings"
)

func minRemoveToMakeValid(s string) string {
	pairCount := getPairCount(s)
	open, closee := pairCount, pairCount

	sb := strings.Builder{}
	for _, r := range s {
		if r == '(' && open > 0 {
            open--
			sb.WriteRune('(')
		} else if r == ')' && closee > open {
            closee--
			sb.WriteRune(')')
		} else if r != '(' && r != ')' {
			sb.WriteRune(r)
		}
	}
    
	return sb.String()
}

func getPairCount(s string) int {
	open, closee := 0, 0

	for _, r := range s {
		if r == '(' {
			open++
		} else if r == ')' && closee < open {
			closee++
		}
	}

	if open < closee {
		return open
	}
	return closee
}