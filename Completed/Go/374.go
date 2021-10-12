/** 
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is lower than the guess number
 *			      1 if num is higher than the guess number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
    if n == 1 {
        return 1
    }
    
    l, r := 0, n
    
    for l <= r {
        mid := (l +r)/2
        g := guess(mid)
        if g == 0 {
            return mid
        } else if g == -1 {
            r = mid-1
        } else {
            l = mid+1
        }
    }
    return -1
}
