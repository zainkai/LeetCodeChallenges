import "container/list"

type frame struct {
	board [2][3]int
	steps int
}

var ans = [2][3]int{
	{1, 2, 3},
	{4, 5, 0},
}

// O(6) == O(1)
func getTypedBoard(board [][]int) [2][3]int {
	typed := [2][3]int{}
	for i, row := range board {
		for j, ele := range row {
			typed[i][j] = ele
		}
	}

	return typed
}

// O(4 * 6 + 6) == O(1) time/space
func getAdjBoard(board [2][3]int) [][2][3]int {
	//find zero
	x, y := -1, -1
	for i, row := range board {
		for j := range row {
			if board[i][j] == 0 {
				x, y = i, j
			}
		}
	}
	// find valid childBoards
	children := [][2][3]int{}
	for _, move := range [][]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}} {
		a, b := x+move[0], y+move[1]
        if a < 0 || b < 0 || a >= len(board) || b >= len(board[a]) {
			continue
		}

		// initalize copy of board
		copied := board
		copied[x][y], copied[a][b] = copied[a][b], copied[x][y]

		children = append(children, copied)
	}
	return children
}

// O(N) time/space
func slidingPuzzle(board [][]int) int {
	initialBoard := getTypedBoard(board)

	q := NewQueue()
	q.Enqueue(frame{
		initialBoard,
		0,
	})
	visited := map[[2][3]int]bool{}

	for q.Len() > 0 {
		f := q.Dequeue().(frame)
		if visited[f.board] {
			continue
		} else if f.board == ans {
			return f.steps
		}
		visited[f.board] = true

		childBoards := getAdjBoard(f.board)
		for _, board := range childBoards {
			q.Enqueue(frame{
				board,
				f.steps + 1,
			})
		}
	}
	return -1
}


// QUEUE UTILS ----------------------------------------------------
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