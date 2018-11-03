
  
function phoneWord(str, val='', map = {}) {
  const dict = {
    '0':[],
    '1':[],
    '2':['a','b','c'],
    '3':['d','e','f'],
    '4':['g','h','i'],
    '5':['j','k','l'],
    '6':['m','n','o'],
    '7':['p','q','r','s'],
    '8':['t','u','v'],
    '9':['w','x','y','z']
  }

  if(val.length === str.length) {
    return val
  }

  const strArr = [...str]
  const s = strArr[val.length]
  for (let i = 0; i < dict[s].length; i++) {
    const newVal = `${val}${dict[s][i]}`
    if(!map[newVal]) {
        map[newVal] = phoneWord(str,newVal,map)
    }
  }
  return Object.keys(map).filter(e => e.length === str.length)
}
  
console.log(phoneWord('23').length)