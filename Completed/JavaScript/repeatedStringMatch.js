function repeatedStringMatch(A, B) {
  let newA = A
  while(newA.length <= B.length) {
    newA += A
    const isRepeated = newA.indexOf(B)
    if (isRepeated !== -1) return isRepeated + 1
  }
  return -1
}