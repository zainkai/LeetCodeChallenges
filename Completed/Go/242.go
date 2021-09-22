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


// better way

func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    
    m := map[byte]int{}
    for i := range s {
        a := s[i]
        m[a]++
        
        b := t[i]
        m[b]--
    }
    
    for _, v := range m {
        if v != 0 {
            return false
        }
    }
    return true
}
