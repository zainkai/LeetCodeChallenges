function lis(arr) {
  const top = (stk) => stk[stk.length -1]
  let max = 1
  let stk = []
  for(let i = 0; i < arr.length -1 ; i++) {
    stk = []
    for(let j = i +1; j < arr.length -1; j++) {
      if(stk.length === 0 || top(stk) <= arr[j]) {
        stk.push(arr[j])
        max = Math.max(max, stk.length)
        console.log(stk)
      } else {
        break
      }
    }
  }
  return max
}