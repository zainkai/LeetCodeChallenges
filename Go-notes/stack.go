package notes

import "errors"

type stack []interface{}

func (s *stack) Push(v interface{}) {
	*s = append(*s, v)
}

func (s *stack) Pop() (interface{}, error) {
	if s.Len() == 0 {
		return nil, errors.New("empty stack")
	}

	res := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return res, nil
}

func (s *stack) Top() (interface{}, error) {
	if s.Len() == 0 {
		return nil, errors.New("empty stack")
	}

	return (*s)[len(*s)-1], nil
}

func (s *stack) Len() int {
	return len(*s)
}
