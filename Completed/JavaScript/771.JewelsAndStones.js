/**
 * @param {string} J
 * @param {string} S
 * @return {number}
 */
var numJewelsInStones = function(J, S) {
  let map = {}
  J.split('').forEach(j => map[j] = 0)
  for(let i =0; i < S.length;i++) {
      const key = S[i]
      if(map[key] !== undefined) {
         map[key]+=1 
      }
  }
  return Object.values(map).reduce((acc ,m) => acc+m,0)
};