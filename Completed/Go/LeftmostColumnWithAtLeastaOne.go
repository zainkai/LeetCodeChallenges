/**
 * // This is the BinaryMatrix's API interface.
 * // You should not implement it, or speculate about its implementation
 * type BinaryMatrix struct {
 *     Get(int, int) int
 *     Dimensions() []int
 * }
 */

 func leftMostColumnWithOne(binaryMatrix BinaryMatrix) int {
    rows := binaryMatrix.Dimensions()[0]
    cols := binaryMatrix.Dimensions()[1]
    
    x,y := 0, cols-1
    found := false
    for y > 0 && x < rows {
        if binaryMatrix.Get(x, y) == 1 {
            found = true
        }
        
        if binaryMatrix.Get(x, y-1) == 1 {
            y--
        } else {
            x++
        }
    }
    
    if found {
        return y
    }
    return -1
}