func stringShift(s string, shift [][]int) string {
	shiftNum := condenseShifts(shift) % len(s)
	if shiftNum == 0 {
		return s
	} else {
        if shiftNum < 0 {
            shiftNum *= -1 
        } else {
            shiftNum = len(s) - shiftNum
        }
        pre := (s[:shiftNum])
        post := (s[shiftNum:])
		s = post + pre
	}

	return s
}

func condenseShifts(shifts [][]int) int {
	shiftNum := 0
	for _, shift := range shifts {
		if shift[0] == 1 {
			shiftNum += shift[1]
		} else {
			shiftNum -= shift[1]
		}
	}
	return shiftNum
}