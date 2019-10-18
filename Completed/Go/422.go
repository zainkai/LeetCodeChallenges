import "errors"

func validWordSquare(words []string) bool {
	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words[i]); j++ {
			a, err1 := getByte(words, i, j)
			b, err2 := getByte(words, j, i)

			if err1 != nil || err2 != nil || a != b {
				return false
			}
		}
	}

	return true
}

func getByte(matrix []string, row int, col int) (byte, error) {
	if row < 0 || row > len(matrix)-1 {
		return byte(0), errors.New("OUT OF BOUNDS")
	} else if col < 0 || col > len(matrix[row])-1 {
		return byte(0), errors.New("OUT OF BOUNDS")
	}

	return matrix[row][col], nil
}