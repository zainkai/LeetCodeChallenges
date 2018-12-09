/**
 * Definition for read4()
 * 
 * @param {character[]} buf Destination buffer
 * @return {number} The number of characters read
 * read4 = function(buf) {
 *     ...
 * };
 */

/**
 * @param {function} read4()
 * @return {function}
 */
var solution = function(read4) {
  /**
   * @param {character[]} buf Destination buffer
   * @param {number} n Maximum number of characters to read
   * @return {number} The number of characters read
   */
  return function(buf, n) {
      let buffer = []
      let total = 0
      while(total < n) {
          read4(buffer)
          for(let i = 0; i < 4; i++) {
              if(total >= n) break
              const char = buffer.shift()
              buf.push(char)
              total +=1
          }
      }
      return total
  }
}