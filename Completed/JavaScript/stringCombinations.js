function stringCombo (str) {
  const map = {}
  ComboHelper(map, str)
  return Object.keys(map)
}

function ComboHelper(map, str) {
  for(let i = 0; i < str.length; i++) {
    const strCopy = Array.from(str)
    const strCopy2 = Array.from(str)
    if(map[str]) {
      return
    }
    map[str] = true
    strCopy.splice(0,1)
    strCopy2.splice(-1,1)
    ComboHelper(map, strCopy.join(""))
    ComboHelper(map, strCopy2.join(""))
  }
}

console.log(stringCombo("abc"))