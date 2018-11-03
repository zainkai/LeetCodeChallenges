function jumpGame(jArr, idx = 0, map = {}) {
  if(idx === jArr.length -1) {
    return true
  }
  if (idx > jArr.length -1) {
    return false
  }

  const maxJump = jArr[idx]
  let result = false
  for(let jump = 1; jump <= maxJump; jump++) {
    if(!map[idx]) {
      map[idx] = jumpGame(jArr, idx + jump, map)
    }
    result = result || map[idx]
  }
  return result
}

console.log(jumpGame([3,2,1,0,4]))