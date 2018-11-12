function kDistictSubStr(str, k) {
  let max = 0
  for(let i = 0 ;i < str.length ; i++) {
    let map = {}
    let count = 0
    for(let j = i;j < str.length;j++) {
      const key = str[j]
      if(map[key] === undefined) {
        map[key] = 1
      } else {
        map[key] += 1
      }

      if(Object.keys(map).length > k) break
      count += 1
      max = Math.max(max, count)
    }
  }

  return max
}


console.log(kDistictSubStr("eceba", 2))