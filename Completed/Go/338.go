// Brute force  O(N * len(int))
func countBits(num int) []int {
	result := make([]int, num+1)
	for i := 0; i <= num; i++ {
		bitsCounted := bit1Count(i)
		result[i] = bitsCounted
	}

	return result
}

func bit1Count(b int, memo *[]int) int {
	numBits := 0
	mask := 1

	for mask <= b {
		if b&mask >= 1 {
			numBits += 1
		}

		mask <<= 1
	}
	return numBits
}

// Optimal O(N)
func countBits(num int) []int {
	result := make([]int, num+1)

	for i := 1; i <= num; i++ {
		halfI := i >> 1  // divide by 2
		if (i & 1) > 0 { // check if last bit is set
			result[i] = result[halfI] + 1
		} else {
			result[i] = result[halfI]
		}
	}

	return result
}