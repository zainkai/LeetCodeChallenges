func isStrobogrammatic(num string) bool {
	m := map[byte]byte{
		'1': '1',
		'6': '9',
		'8': '8',
		'9': '6',
		'0': '0',
	}

	i, j := 0, len(num)-1
	for i <= j {
		a := num[i]
		b := num[j]
		if m[a] != b {
			return false
		}

		i++
		j--
	}
	return true
}