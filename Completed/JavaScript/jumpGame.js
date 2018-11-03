function jumpGame(jArr, idx = 0) {
  if(idx === jArr.length -1) {
    return true
  }
  if (idx > jArr.length -1) {
    return false
  }

  const maxJump = jArr[idx]
  let result = false
  for(let jump = 1; jump <= maxJump; jump++) {
    result = result || jumpGame(jArr, idx + jump)
  }

  return result
}

console.log(jumpGame([3,2,1,0,4]))