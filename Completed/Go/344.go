package LeetCode

func reverseString(s []byte) {
	l := int(len(s) / 2)
	for i := 0; i < l; i++ {
		rIdx := len(s) - i - 1
		temp := s[i]

		s[i] = s[rIdx]
		s[rIdx] = temp
	}
}
