/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 import "container/list"

 type frame struct {
	 node     *TreeNode
	 vertical int
 }
 
 func verticalOrder(root *TreeNode) [][]int {
	 q := NewQueue()
	 if root != nil {
		 q.Enqueue(frame{root, 0})
	 } else {
		 return [][]int{}
	 }
 
	 vertMap := map[int][]int{}
	 leftMost, rightMost := 0, 0
 
	 for q.Len() > 0 {
		 f := q.Dequeue().(frame)
		 vertMap[f.vertical] = append(vertMap[f.vertical], f.node.Val)
 
		 if f.vertical < leftMost {
			 leftMost = f.vertical
		 } else if f.vertical > rightMost {
			 rightMost = f.vertical
		 }
 
		 if f.node.Left != nil {
			 q.Enqueue(frame{f.node.Left, f.vertical - 1})
		 }
		 if f.node.Right != nil {
			 q.Enqueue(frame{f.node.Right, f.vertical + 1})
		 }
	 }
 
	 ans := [][]int{}
	 for i := leftMost; i <= rightMost; i++ {
		 ans = append(ans, vertMap[i])
	 }
	 return ans
 }
 
 // Queue ...
 type Queue struct {
	 list *list.List
 }
 
 // NewQueue ...
 func NewQueue() *Queue {
	 return &Queue{
		 list: list.New(),
	 }
 }
 
 // Len ...
 func (q *Queue) Len() int {
	 return q.list.Len()
 }
 
 // Enqueue ...
 func (q *Queue) Enqueue(val interface{}) {
	 q.list.PushBack(val)
 }
 
 // Dequeue ...
 func (q *Queue) Dequeue() interface{} {
 
	 ele := q.list.Front()
	 defer q.list.Remove(ele)
 
	 return ele.Value
 }
 
 // Top ...
 func (q *Queue) Top() interface{} {
	 ele := q.list.Front()
	 return ele.Value
 }