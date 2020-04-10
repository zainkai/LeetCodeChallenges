/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (this NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (this NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (this *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (this NestedInteger) GetList() []*NestedInteger {}
 */

 type NestedIterator struct {
	stk ElementStk
}

type Element struct {
	data []*NestedInteger
	idx    int
}

type ElementStk []Element

func (e *ElementStk) Push(x Element) {
	*e = append(*e, x)
}

func (e *ElementStk) Len() int {
	return len(*e)
}

func (e *ElementStk) Top() Element {
    if len(*e) == 0 {
        return Element{}
    }
    
	return (*e)[len(*e)-1]
}

func (e *ElementStk) Pop() Element {
	top := e.Top()
	*e = (*e)[:len(*e)-1]

	return top
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	construct := &NestedIterator{
		stk: make(ElementStk, 0),
	}

	if len(nestedList) == 0 {
		return construct
	}
	construct.stk.Push(Element{nestedList, 0})

	return construct
}

func (this *NestedIterator) Next() int {
	top := this.stk.Pop()

	if top.data[top.idx].IsInteger() {
		if top.idx < len(top.data)-1 {
			this.stk.Push(Element{top.data, top.idx + 1})
		}

		return top.data[top.idx].GetInteger()
	} else {
        if top.idx < len(top.data)-1 {
			this.stk.Push(Element{top.data, top.idx + 1})
		}

		nextList := top.data[top.idx].GetList()
		this.stk.Push(Element{nextList, 0})

		return this.Next()
	}
}

func (this *NestedIterator) HasNext() bool {
    top := this.stk.Top()
    for this.stk.Len() > 0 && !top.data[top.idx].IsInteger() {
        top = this.stk.Pop()
        if top.idx < len(top.data)-1 {
			this.stk.Push(Element{top.data, top.idx + 1})
		}
        nextList := top.data[top.idx].GetList()
        
        if len(nextList) != 0 {
           this.stk.Push(Element{nextList, 0}) 
        }
        top = this.stk.Top()
    }
    
    
	return this.stk.Len() > 0
}