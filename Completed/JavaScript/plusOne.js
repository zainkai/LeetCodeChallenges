function plusOne(num) {
  if(num.length === 0) return []
  let carry = 1
  
  for(let i = num.length -1; i >= 0; i--) {
    num[i] += carry
    if(num[i] > 9) {
      carry = 1
      num[i] = num[i] % 10
    } else {
      carry = 0
      break
    }
  }

  if(carry !== 0) {
    num.unshift(carry)
  }

  return num
}