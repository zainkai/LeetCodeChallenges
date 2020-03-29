package notes

import "container/list"

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
func (q *Queue) Dequeue() (interface, error) {
	if q.Len() == 0 {
		return nil, errors.New("empty queue")
	}

	ele := q.list.Front()
	defer q.list.Remove(ele)

	return ele.Value, nil
}

// Top ...
func (q *Queue) Top() (interface, error) {
	if q.Len() == 0 {
		return nil, errors.New("empty queue")
	}

	ele := q.list.Front()
	return ele.Value, nil
}
