/**
 * @param {number[][]} image
 * @param {number} sr
 * @param {number} sc
 * @param {number} newColor
 * @return {number[][]}
 */
var floodFill = function(image, sr, sc, newColor) {
  const target = image[sr][sc]
  if (target === newColor) return image
  
  
  const queue = [[sr,sc]]
  while(queue.length > 0) {
      const val = queue.shift()
      const r = val[0]
      const c = val[1]
      
      if(image[r] === undefined || image[r][c] === undefined || image[r][c] !== target) continue
      
      image[r][c] = newColor
      queue.push([r+1, c])
      queue.push([r-1, c])
      queue.push([r, c+1])
      queue.push([r, c-1])
  }
  return image
}