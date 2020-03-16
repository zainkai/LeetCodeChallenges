import (
	"errors"
)

func validateStackSequences(pushed []int, popped []int) bool {
	stk := stack{}
	pu, po := 0, 0

	for pu <= len(pushed) && po < len(popped) {
		if t, err := stk.Top(); err == nil && t == popped[po] {
			stk.Pop()
			po++
		} else if pu < len(pushed) {
			stk.Push(pushed[pu])
			pu++
		} else {
			return false
		}
	}

	return true
}

// stack utils
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