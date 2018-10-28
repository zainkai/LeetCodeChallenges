function flipAndInvertImage (A) {
  const res = []
  A.forEach(bitArr => {
    bitArr = bitArr.reverse().map(bit => bit === 1 ? 0 : 1)
    res.push(bitArr)
  })

  return res
}

function flipAndInvertImage2 (A) {
  return A.map(bitArr => bitArr.reverse().map(bit => bit === 1 ? 0 : 1))
}

const result = [[1,1,0],[1,0,1],[0,0,0]]
console.log(flipAndInvertImage2(result))