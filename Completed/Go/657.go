var moveMap = map[rune][]int{
	'U': {0, 1},
	'D': {0, -1},
	'L': {-1, 0},
	'R': {1, 0},
}

func judgeCircle(moves string) bool {
	posX, posY := 0, 0
	for _, m := range moves {
		curr := moveMap[m]

		posX += curr[0]
		posY += curr[1]
	}

	return posX == 0 && posY == 0
}