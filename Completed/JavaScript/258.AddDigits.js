/**
 * @param {number} num
 * @return {number}
 */
var addDigits = function(num) {
  let arr = num.toString().split('')
  while(arr.length !== 1) {
    arr = arr
      .map(a => Number(a))
      .reduce((acc,a) => a+acc,0)
      .toString()
      .split('')
  }
  return arr[0]
}