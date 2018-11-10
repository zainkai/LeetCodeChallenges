function zigzagConversion (str, rows) {
  if(rows === 0) return ""
  if(rows === 1) return str

  let resArr = new Array(rows)
  for(let i = 0; i < rows; i++) {
    resArr[i] = []
  }

  let isDown = true
  let rowArr = 0
  for(let i = 0;i < str.length; i++) {
    if(isDown) {
      resArr[rowArr].push(str[i])
      rowArr++
    } else {
      resArr[rowArr].push(str[i])
      rowArr--
    }
    isDown = rowArr === rows -1 || rowArr === 0 ? !isDown : isDown
  }
  
  return resArr.reduce((acc,cur) => acc += cur.join(""), "")
}


console.log(zigzagConversion('PAYPALISHIRING', 3))