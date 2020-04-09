func backspaceCompare(S string, T string) bool {
    sS, sT := runeStk{}, runeStk{}
    for _, r := range S {
        if r == '#' {
            sS.Pop()
        } else {
            sS.Push(r)
        }
    }
    
    for _, r := range T {
        if r == '#' {
            sT.Pop()
        } else {
            sT.Push(r)
        }
    }
    
    return sS.Keys() == sT.Keys()
}

type runeStk []rune

func (r *runeStk) Push(x rune) {
	*r = append(*r, x)
}

func (r *runeStk) Pop() rune {
	if len(*r) == 0 {
		return rune(0)
	}

	top := (*r)[len(*r)-1]
	*r = (*r)[:len(*r)-1]

	return top
}

func (r *runeStk) Keys() string {
	return string(*r)
}