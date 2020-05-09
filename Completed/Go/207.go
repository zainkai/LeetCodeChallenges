import "container/list"

func canFinish(numCourses int, prerequisites [][]int) bool {
    visited := map[int]bool{}
    indegree := map[int]int{}
    
    for _, edge := range prerequisites {
        next := edge[0]
        indegree[next]++
    }
    
    q := NewQueue()
    for i := 0; i < numCourses; i++ {
        if indegree[i] == 0 {
            q.Enqueue(i)
        }
    }
    
    for q.Len() > 0 {
        course := q.Dequeue().(int)
        visited[course] = true
        for _, edge := range prerequisites {
            if edge[1] == course {
                indegree[edge[0]]--
            }
        }
        QueueZeroDegree(indegree, visited, q)
    }
    
    return len(indegree) == 0
}

func QueueZeroDegree(indegree map[int]int, visited map[int]bool, q *Queue) {
    for course, deg := range indegree {
        if deg == 0 && !visited[course] {
            q.Enqueue(course)
            delete(indegree, course)
        }
    }
}


// Queue util................................................
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
func (q *Queue) Dequeue() (interface{}) {
	ele := q.list.Front()
	defer q.list.Remove(ele)

	return ele.Value
}

// Top ...
func (q *Queue) Top() (interface{}) {
	ele := q.list.Front()
	return ele.Value
}
