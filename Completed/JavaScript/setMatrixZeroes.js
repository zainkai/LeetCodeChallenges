//row - n
//col - m
// O(1) - space

function reset(matrix,row,col, value) {
  for(let i = 0; i < matrix[row].length ; i++) { //O(m)
    if(matrix[row][i] === 0) continue
    matrix[row][i] = value
  }
  for(let i = 0; i < matrix.length; i++) { //O(n)
    if(matrix[i][col] === 0) continue
    matrix[i][col] = value
  }
}




function setMatrixZeros(matrix) {
  for(let i = 0; i < matrix.length; i++) { //O(nm)
    for(let j = 0; j < matrix[i].length; j++) {
      if(matrix[i][j] === 0) {
        reset(matrix,i,j,Infinity)
      }
    }
  }

  for(let i = 0; i < matrix.length; i++) {//O(nm)
    for(let j = 0; j < matrix[i].length; j++) {
      if(matrix[i][j] === Infinity) {
        matrix[i][j] = 0
      }
    }
  }
}
