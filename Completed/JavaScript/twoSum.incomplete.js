function twoSum (arr, target) {
  const compMap = []
  let result = []
  for(let num of arr) {
    const complement = target - num
    const foundVal = compMap.find(c =>  c === num)
    if (foundVal) {
      return result = [complement, num]
    } else {
      compMap.push(complement)
    }
  }
  