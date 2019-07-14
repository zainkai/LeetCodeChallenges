/**
 * @param {string} s
 * @return {string}
 */
var longestPalindrome = function(s) {
  if(!s || !s.length || typeof s !== 'string') return ''
  
  let result = ''
  for(let i = 0; i < s.length;i++) {
      const t1 = findPalin(s, i,i) // odd length palin
      const t2 = findPalin(s, i, i+1) // even length palin
      
      if (t1.length >= result.length) {
          result = t1
      }
      
      if (t2.length >= result.length) {
          result = t2
      }
  }
  return result
};

function findPalin(s, l, r) {
let result = ''
if(s[l] !== s[r]) { return result }

while(s[l] !== undefined && s[l] === s[r]) {
  result = s.slice(l, r+1)
  l -= 1
  r += 1
}

return result
}