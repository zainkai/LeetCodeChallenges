import "container/list"

type QueueNode struct {
	n *TreeNode
	level int
}


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
    if root == nil {
        return make([][]int, 0)
    }
    
    q := NewQueue()

	res := make([][]int, 0)
	q.Enqueue(QueueNode{root, 1})
	
	for q.Len() > 0 {
		x, _ := q.Dequeue()
        qn := x.(QueueNode)
        if qn.level > len(res) {
            res = append(res, []int{})
        }


        res[qn.level-1] = append(res[qn.level-1], qn.n.Val)

        if qn.n.Left != nil {
            q.Enqueue(QueueNode{ qn.n.Left, qn.level+1 })
        }
        if qn.n.Right != nil {
            q.Enqueue(QueueNode{ qn.n.Right, qn.level+1 })
        }

    }
    return res
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
func (q *Queue) Dequeue() (interface{}, error) {
	if q.Len() == 0 {
		return nil, errors.New("empty queue")
	}

	ele := q.list.Front()
	defer q.list.Remove(ele)

	return ele.Value, nil
}

// Top ...
func (q *Queue) Top() (interface{}, error) {
	if q.Len() == 0 {
		return nil, errors.New("empty queue")
	}

	ele := q.list.Front()
	return ele.Value, nil
}