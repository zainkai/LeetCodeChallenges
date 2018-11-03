const uniquePaths = (m,n, rIdx = 0, cIdx = 0, map={}) => {
  if(rIdx === m -1 && cIdx === n -1) return 1
  if(rIdx > m -1) return 0
  if(cIdx > n -1) return 0


  const hashKeyDown = [rIdx+1,cIdx].toString() // 0,0
  if(!map[hashKeyDown]){
    map[hashKeyDown] = uniquePaths(m,n, rIdx +1, cIdx, map)
  }
  const hashKeyRight = [rIdx,cIdx +1].toString()
  if(!map[hashKeyRight]){
    map[hashKeyRight] = uniquePaths(m,n, rIdx, cIdx+ 1, map)
  }
  
  return (
    map[hashKeyDown]
    +
    map[hashKeyRight]
  )
}

console.log(uniquePaths(7,3))