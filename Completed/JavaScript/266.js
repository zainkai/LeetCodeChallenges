/**
 * @param {string} s
 * @return {boolean}
 */
var canPermutePalindrome = function(s) {
  if (typeof s !== 'string') return false
  if (s.length === 1 || s.length === 0) return true
  
  const map = {}
  
  for (let char of s) {
      if (map[char] !== undefined) {
          map[char] += 1
      } else {
          map[char] = 1
      }
  }
  
  let oddCount = 0
  for (let key in map) {
      const value = map[key]
      if (map[key] % 2 !== 0) { // odd
          oddCount += 1
      }
      
      if (oddCount > 1) return false
  }
  
  return true
};