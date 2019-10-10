type runeStack struct {
	data []rune
}

func isValid(s string) bool {
	stk := runeStack{}
	for _, r := range s {
		if top := stk.Peek(); isPair(top, r) {
			stk.Pop()
		} else {
			stk.Push(r)
		}
	}

	return stk.Len() == 0
}

func (this *runeStack) Peek() rune {
	var top rune
	if len(this.data) != 0 {
		top = this.data[len(this.data)-1]
	}
	return top
}

func (this *runeStack) Push(x rune) {
	this.data = append(this.data, x)
}

func (this *runeStack) Pop() rune {
	x := this.Peek()
	this.data = this.data[:len(this.data)-1]

	return x
}

func (this *runeStack) Len() int {
	return len(this.data)
}

func isPair(a, b rune) bool {
	m := map[rune]rune{
		'(': ')', '[': ']', '{': '}',
	}
	val, ok := m[a]
	return ok && val == b
}