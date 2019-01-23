/**
 * @param {number[]} A
 * @return {number[]}
 */
var sortArrayByParity = function(A) {
  const even = [], odd = []
  A.forEach(a => {
      if (a % 2 === 0 ) even.push(a)
      else odd.push(a)
  })
  return [...even, ...odd]  
};