/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 type BSTIterator struct {
    stk stack
}


func Constructor(root *TreeNode) BSTIterator {
    iter := BSTIterator{}
    iter.stk.AddNode(root)
    
    return iter
}


/** @return the next smallest number */
func (this *BSTIterator) Next() int {
    n, _ := this.stk.Pop()
    if n.Right != nil {
        this.stk.AddNode(n.Right)
    }
    return n.Val
}


/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
    return this.stk.Len() > 0
}


/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */

type stack []*TreeNode

func (s *stack) Push(v *TreeNode) {
	*s = append(*s, v)
}

func (s *stack) AddNode(v *TreeNode) {
    curr := v
    for curr != nil {
        s.Push(curr)
        curr = curr.Left
    }
}

func (s *stack) Pop() (*TreeNode, error) {
	if s.Len() == 0 {
		return nil, errors.New("empty stack")
	}

	res := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return res, nil
}

func (s *stack) Top() (*TreeNode, error) {
	if s.Len() == 0 {
		return nil, errors.New("empty stack")
	}

	return (*s)[len(*s)-1], nil
}

func (s *stack) Len() int {
	return len(*s)
}