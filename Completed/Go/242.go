import "sort"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	x := sortBytes(s)
	y := sortBytes(t)

	for i := 0; i < len(s); i++ {
		if x[i] != y[i] {
			return false
		}
	}

	return true
}

func sortBytes(s string) []byte {
	arr := []byte(s)

	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	return arr
}