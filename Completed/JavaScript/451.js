/**
 * @param {string} s
 * @return {string}
 */
var frequencySort = function(s) {
  // counting
  
  const map = {}
  for (let char of s) {
      map[char] = (map[char] || 0) + 1
  }
  
  return Object.keys(map).sort((a,b) => map[b] - map[a]).reduce((str, key) => {
      for(let i = 0; i < map[key]; i++) {
          str += key
      }
      
      return str
  }, '')
};