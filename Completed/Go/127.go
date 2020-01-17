type item struct {
	Word string
	Dist int
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	dict, q, result := getOneDistDictionary(beginWord, wordList), New(), 0
	q.Enqueue(&item{beginWord, 0})

	visited := map[string]bool{}
	for q.Len() > 0 {
		ele := q.Dequeue().(*item)

		if ele.Word == endWord && (result == 0 || ele.Dist < result) {
			result = ele.Dist + 1
		}

		for _, word := range dict[ele.Word] {
			if visited[word] {
				continue
			}
			q.Enqueue(&item{word, ele.Dist + 1})
			visited[word] = true
		}
	}

	return result
}

func isOneEditDist(s, t string) bool {
	diffFound := false
	for i := 0; i < len(s); i++ {
		if s[i] == t[i] {
			continue
		} else if s[i] != t[i] && diffFound {
			return false
		} else {
			diffFound = true
		}
	}
	return true
}

// getOneDistDictionary get all words with a edit distance of one
func getOneDistDictionary(beginWord string, wordList []string) map[string][]string {
	dict := make(map[string][]string)
	wordList = append(wordList, beginWord)

	for i, key := range wordList {
		for j, word := range wordList {
			if i == j {
				continue
			} else if isOneEditDist(key, word) {
				dict[key] = append(dict[key], word)
			}
		}
	}

	return dict
}

//https://github.com/zainkai/go-collections/blob/master/queue/queue.go
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