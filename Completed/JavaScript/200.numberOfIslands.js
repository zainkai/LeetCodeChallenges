function fillIsland(grid, row, col) {
  if(row > grid.length -1 || row < 0) return
  if(col > grid[row].length -1 || col < 0) return
  if(grid[row][col] === "0") return
  if(grid[row][col] === "1") grid[row][col] = "0"

  fillIsland(grid, row -1, col)
  fillIsland(grid, row +1, col)
  fillIsland(grid, row, col -1)
  fillIsland(grid, row, col +1)
}

function numberOfIslands(grid) {
  let numIslands = 0
  for(let i = 0; i < grid.length; i++) {
    for(let j = 0; j < grid[i].length; j++) {
      if(grid[i][j] === "1") {
        numIslands +=1
        
        fillIsland(grid,i,j)
      }
    }
  }
  return numIslands
}

let grid = [["1","1","1","1","0"],["1","1","0","1","0"],["1","1","0","0","0"],["0","0","0","0","0"]]
console.log(numberOfIslands(grid))