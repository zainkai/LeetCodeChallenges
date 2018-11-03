function coinChange(coinArr, total, acc = 0, numCoin = 0, map = {}) {
  if(acc > total) return Infinity
  if(acc === total) return numCoin

  let result = Infinity
  for(let coin of coinArr) {
    if(!map[acc + coin]) {
      map[acc + coin] = coinChange(coinArr, total, acc + coin, numCoin +1)
    }

    result = Math.min(
      result,
      map[acc + coin]
    )
  }
  return result === Infinity ? -1 : result
}

console.log(coinChange([1,2,5], 11))