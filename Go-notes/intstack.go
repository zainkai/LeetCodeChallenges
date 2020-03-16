package notes

import "errors"

type stack []int

func (s *stack) Push(v int) {
	*s = append(*s, v)
}

func (s *stack) Pop() (int, error) {
	if s.Len() == 0 {
		return 0, errors.New("Empty Stack")
	}

	res := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return res, nil
}

func (s *stack) Top() (int, error) {
	if s.Len() == 0 {
		return 0, errors.New("Empty Stack")
	}

	return (*s)[len(*s)-1], nil
}

func (s *stack) Len() int {
	return len(*s)
}
