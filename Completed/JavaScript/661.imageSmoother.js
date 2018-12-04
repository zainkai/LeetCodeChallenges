function imageSmoother(matrix) {
  if(matrix.length === 0) return []
  let cArr = new Array(matrix.length)
  matrix.forEach((a,i) => cArr[i] = new Array(a.length).fill(null))

  for(let i = 0; i < matrix.length; i++) {
    for(let j = 0; j < matrix[i].length;j++) {
      cArr[i][j] = smooth(matrix,i,j)
    }
  }

  return cArr
}

function smooth(matrix, x ,y) {
  let result = 0
  let cells = 0
  for (let i = x -1; i < x+ 2; i++) {
    for(let j = y-1; j < y +2; j++) {
      if(matrix[i] !== undefined && matrix[i][j] !== undefined) {
        cells +=1
        result += matrix[i][j]
      }
    }
  }
  return Math.floor(result / cells)
}

const prob = [
  [1,1,1],
  [1,1,1],
  [1,1,1]
]
console.log(imageSmoother(prob))
