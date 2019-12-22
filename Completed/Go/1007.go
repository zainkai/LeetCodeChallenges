func minDominoRotations(A []int, B []int) int {
	minFlips := (1 << 63) - 1

	if x := helper(A, B, A[0]); x != -1 && x < minFlips {
		minFlips = x
	}

	if x := helper(B, A, A[0]); x != -1 && x < minFlips {
		minFlips = x
	}

	if x := helper(A, B, B[0]); x != -1 && x < minFlips {
		minFlips = x
	}

	if x := helper(B, A, B[0]); x != -1 && x < minFlips {
		minFlips = x
	}

	if minFlips == (1<<63)-1 {
		return -1
	}
	return minFlips
}

func helper(A []int, B []int, d int) int {
	temp := 0

	for i := 0; i < len(A); i++ {
		if A[i] != d && B[i] != d {
			return -1
		} else if A[i] == d {
			continue
		} else if B[i] == d {
			temp++
		}
	}
	return temp
}