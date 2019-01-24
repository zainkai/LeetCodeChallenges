/**
 * @param {string} S
 * @return {number[]}
 */
var diStringMatch = function(S) {
  let low = 0, high = S.length
  let result = []
  S.split('').forEach((s) => {
      if (s === 'I') {
          result.push(low)
          low++
      } else {
          result.push(high)
          high--
      }
  })
  result.push(low)
  return result
};