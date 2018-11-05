function interSect2Arrs(num1, num2) {
  let map1 = {}
  let stk = []

  num1.forEach(e => { //O(n1)
    map1[e] = map1 + 1 || 1
  })

  console.log(map1)
  num2.forEach(e => { //O(n2)
    if(map1[e] !== undefined && map1[e] !== 0) {
      map1[e] = map1[e] -  1
      stk.push(e)
    }
  })

  return stk
}

