// error with [0,0,0]

function(arrNums) {
  let map = {}
  let result = []

  let x, y, z
  const len = arrNums.length -1
  for (x = 0; x < len; x++) {
   for(y = x+1; y < len; y++) {
    for(z = y+1; z < len; z++) {
      const key = [arrNums[x],arrNums[y],arrNums[z]].sort().join()
      console.log(key)
      if (map[key] !== undefined) {continue}
          
      if(!map[key]) {
        map[key] = arrNums[x] + arrNums[y] + arrNums[z]
      }

      if (map[key] === 0) {
        result.push([arrNums[x], arrNums[y], arrNums[z]])
      }
    }
   }
  }
  return result
}