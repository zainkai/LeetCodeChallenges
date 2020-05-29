type Point struct {
    r, c int
}

type SnakeGame struct {
    s []Point
    sLen int
    
    foods []Point
    width int
    height int
}


/** Initialize your data structure here.
        @param width - screen width
        @param height - screen height 
        @param food - A list of food positions
        E.g food = [[1,1], [1,0]] means the first food is positioned at [1,1], the second is at [1,0]. */
func Constructor(width int, height int, food [][]int) SnakeGame {
    foods := make([]Point, len(food))
    for i, edge := range food {
        foods[i] = Point{edge[0], edge[1]} 
    }
    
    return SnakeGame{
        s: []Point{{0,0}},
        sLen: 1,
        foods: foods,
        width: width,
        height: height,
    }
}


/** Moves the snake.
        @param direction - 'U' = Up, 'L' = Left, 'R' = Right, 'D' = Down 
        @return The game's score after the move. Return -1 if game over. 
        Game over when snake crosses the screen boundary or bites its body. */
func (this *SnakeGame) Move(direction string) int {
    cpyHead := this.s[len(this.s)-1]
    switch direction {
    case "U":
        cpyHead.r--
    case "D":
        cpyHead.r++
    case "L":
        cpyHead.c--
    case "R":
        cpyHead.c++
    }
    
    if cpyHead.r < 0 || cpyHead.c < 0 {
        return -1
    } else if cpyHead.r >= this.height || cpyHead.c >= this.width {
        return -1
    }
    
    if this.checkbiteSelf(cpyHead) {
        return -1
    }
    this.checkFood(cpyHead)
    return this.sLen -1
}

func (this *SnakeGame) checkbiteSelf(head Point) bool {
    for i := 1; i < len(this.s); i++ { // ok to bite tail
        if this.s[i] == head {
            return true
        }
    }
    return false
}

func (this *SnakeGame) checkFood(head Point) {
    if len(this.foods) > 0 && this.foods[0] == head {
        this.sLen++
        this.foods = this.foods[1:]
    }
    
    this.s = append(this.s, head)
    if len(this.s) > this.sLen {
        this.s = this.s[1:]
    }
}


/**
 * Your SnakeGame object will be instantiated and called as such:
 * obj := Constructor(width, height, food);
 * param_1 := obj.Move(direction);
 */
