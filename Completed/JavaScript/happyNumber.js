function happyNumber(n) {
  let sum = n

  while(sum !== 1) {
    sum = [...(sum.toString())]
      .map(e => Number(e))
      .reduce((acc, cur) => {
      acc += Math.pow(cur, 2)
      return acc
    }, 0)
  }

  return true
}

console.log(happyNumber(19))