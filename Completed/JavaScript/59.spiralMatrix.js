function rotateCounterClockWise (matrix) {
  if(matrix.length === 0) return []
  const rows = matrix.length, cols = matrix[0].length

  let newMatrix = new Array(cols)
  for(let i = 0; i < cols; i++) {
    newMatrix[i] = new Array(rows)
  }
  
  for(let i = 0; i < rows; i++) {
    for(let j = 0; j < cols; j++) {
      newMatrix[j][i] = matrix[i][j]
    }
  }
  return newMatrix.reverse()
}

function spiralMatrix(matrix) {
  let stk = []
  let matrixPtr = matrix
  while(matrixPtr.length !== 0) {
    for(let i =0; i< matrixPtr[0].length; i++) {
      stk.push(matrixPtr[0][i])
    }
    matrixPtr.shift()
    matrixPtr = rotateCounterClockWise(matrixPtr)
  }
  return stk
}


let arr = [
  [1, 2, 3, 4],
  [5, 6, 7, 8],
  [9,10,11,12]
]

console.log(spiralMatrix(arr))
