/*   Below is the interface for Iterator, which is already defined for you.
 *
 *   type Iterator struct {
 *       
 *   }
 *
 *   func (this *Iterator) hasNext() bool {
 *		// Returns true if the iteration has more elements.
 *   }
 *
 *   func (this *Iterator) next() int {
 *		// Returns the next element in the iteration.
 *   }
 */

type PeekingIterator struct {
    q []int
    iter *Iterator
}

func Constructor(iter *Iterator) *PeekingIterator {
    return &PeekingIterator{
        make([]int, 0),
        iter,
    }
}

func (this *PeekingIterator) hasNext() bool {
    if len(this.q) > 0 || this.iter.hasNext() {
        return true
    }
    return false
}

func (this *PeekingIterator) next() int {
    if len(this.q) > 0 {
        top := this.q[0]
        this.q = this.q[1:]
        
        return top
    }
    return this.iter.next()
}

func (this *PeekingIterator) peek() int {
    if len(this.q) == 0 {
        this.q = append(this.q, this.iter.next())
    }
    
    return this.q[0]
}
