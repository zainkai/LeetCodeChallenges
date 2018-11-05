function lengthOfLongestSubstring (str) {
  if(str.length === 1) return 1
  let result = ""
  let stk = []
  for(let i = 0; i < str.length; i++) {
    for(let j = i; j < str.length; j++) {
      let s = str[j]
      if(stk.findIndex(e => e === s) === -1) {
        stk.push(s)
      } else {
        result = result.length > stk.length ? result : stk.join("")
        stk = []
        break
      }
    }
  }
  
  result = result.length > stk.length ? result : stk.join("")
  return result.length
}


console.log(lengthOfLongestSubstring("abcabcbb"))