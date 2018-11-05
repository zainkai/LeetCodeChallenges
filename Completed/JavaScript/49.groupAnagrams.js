function groupAnagrams(strs) {
  const map = {}

  strs.forEach(e => {
    const key = [...e].sort().join()
    if(map[key] === undefined) {
      map[key] = []
    }
    if(map[key] !== undefined) {
      map[key].push(e)
    }
  })

  return Object.keys(map).map(key => map[key])
}



