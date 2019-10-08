func isOneBitCharacter(bits []int) bool {
	i := 0
	for i < len(bits)-2 {
		bit := bits[i]
		if bit == 1 {
			i += 2
		} else if bit == 0 {
			i += 1
		}
	}

	return bits[i] == 0
}