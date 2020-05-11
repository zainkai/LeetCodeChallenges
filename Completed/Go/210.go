func findOrder(numCourses int, prereqs [][]int) []int {
    indegree := map[int]int{}
    visited := map[int]bool{}
    adjList := map[int][]int{}
    q := NewQueue()
    
    for i := 0; i < numCourses; i++ {
        indegree[i] = 0
    }
    for _, prereq := range prereqs {
        start, end := prereq[1], prereq[0]
        indegree[end]++
        
        adjList[start] = append(adjList[start], end)
    }
    
    path := []int{}
    QueueZeroDegree(indegree, visited, q, &path)
    
    for q.Len() > 0 {
        course := q.Dequeue().(int)
        visited[course] = true
        
        for _, next := range adjList[course] {
            indegree[next]--
        }
        QueueZeroDegree(indegree, visited, q, &path)
    }
    
    if len(indegree) != 0 {
        return []int{}
    }
    return path
}

func QueueZeroDegree(indegree map[int]int, visited map[int]bool, q *Queue, path *[]int) {
    for course, deg := range indegree {
        if deg == 0 && !visited[course] {
            *path = append(*path, course)
            
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
