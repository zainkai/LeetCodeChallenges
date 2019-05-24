/**
 * @param {number[][]} grid
 * @return {number}
 */
var maxIncreaseKeepingSkyline = function(grid) {
  let tb = new Array(grid.length).fill(0)
  let lr = new Array(grid.length).fill(0)
  
  for(let a = 0; a < grid.length;a++) {
      for(let b = 0; b < grid[a].length;b++) {
          tb[a] = Math.max(tb[a], grid[a][b])
          lr[b] = Math.max(lr[b], grid[a][b])
      }
  }
  let res = 0
  for(let a = 0; a < grid.length;a++) {
      for(let b = 0; b < grid[a].length;b++) {
          let diff = Math.min(tb[a], lr[b])
          res += diff - grid[a][b]
      }
  }
  return res
};