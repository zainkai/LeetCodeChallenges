type Move struct {
	X, Y  int
	Jumps int
}

type coordinate struct {
	X, Y int
}

func minKnightMoves(x int, y int) int {
	q := New()
	q.Enqueue(Move{0, 0, 0})

	memo := map[coordinate]bool{}

	for q.Len() > 0 {
		mov := q.Dequeue().(Move)
		coor := coordinate{mov.X, mov.Y}

		if memo[coor] {
			continue
		} else {
			memo[coor] = true
		}

		if mov.X == x && mov.Y == y {
			return mov.Jumps
		}
		addNextMoves(q, mov)
	}

	return -1
}

func addNextMoves(q *Queue, currMov Move) {
	directions := [][]int{[]int{-2, -1}, []int{-2, 1}, []int{-1, 2}, []int{1, 2}, []int{2, 1}, []int{2, -1}, []int{1, -2}, []int{-1, -2}}
	for _, dirs := range directions {
		q.Enqueue(Move{currMov.X + dirs[0], currMov.Y + dirs[1], currMov.Jumps + 1})
	}
}

///-------------------------------------- utils
type QueueNode struct {
	Value interface{}
	Next  *QueueNode
}

// Queue generic stack implementation using a Singly Linked List
type Queue struct {
	Head   *QueueNode
	Tail   *QueueNode
	Length int
}

// New Create a new queue
// O(1)
func New() *Queue {
	return &Queue{
		Head:   nil,
		Tail:   nil,
		Length: 0,
	}
}

// Dequeue Take the next item off the front of the queue
// O(1)
func (q *Queue) Dequeue() interface{} {
	if q.Head == nil {
		return nil
	}

	front := q.Peek()
	if q.Head == q.Tail {
		q.Head = nil
		q.Tail = nil
	} else {
		q.Head = q.Head.Next
	}

	q.Length--
	return front
}

// Enqueue Put an item on the end of a queue
// O(1)
func (q *Queue) Enqueue(value interface{}) {
	t := &QueueNode{
		Next:  nil,
		Value: value,
	}

	if q.Head == nil {
		q.Head = t
		q.Tail = t
	} else {
		q.Tail.Next = t
		q.Tail = t
	}

	q.Length++
}

// Len Return the number of items in the queue
// O(1)
func (q *Queue) Len() int {
	return q.Length
}

// Peek Return the first item in the queue without removing it
// O(1)
func (q *Queue) Peek() interface{} {
	if q.Head == nil {
		return nil
	}

	return q.Head.Value
}