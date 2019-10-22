import "strings"

func repeatedStringMatch(A string, B string) int {
	t := strings.Builder{}
	t.WriteString(A)

	i := 1
	l := len(B)/len(A) + 2
	for i <= l {
		if strings.Index(t.String(), B) != -1 {
			return i
		}
		t.WriteString(A)
		i++
	}
	return -1
}