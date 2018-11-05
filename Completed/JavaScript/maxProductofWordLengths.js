function isWordIntersect(w1, w2) {
  const s1 = [...w1]
  const s2 = [...w2]
  for(let i = 0 ; i < s1.length; i++) {
    if(s2.findIndex(e => e === s1[i]) !== -1) return true
  }
  return false
}

function maxProduct(words) {
  let product = 0
  for(let i = 0; i < words.length; i++) {
    for(let j = 0; j < words.length; j++) {
      if(i == j) continue
      if(!isWordIntersect(words[i], words[j])) {
        const value = words[i].length * words[j].length
        product = Math.max(product, value)
      }
    }
  }
  return product
}
