function KeysAndRooms(rooms) {
  let visited = new Array(rooms.length).fill(false)
  visited[0] = true
  let q = [0]
  while(q.length !== 0) {
      const s = q.shift()
      rooms[s].forEach(key => {
          if(!visited[key]) {
              q.push(key)
              visited[key] = true
          }
      })
  }
  return visited.reduce((acc,v) => acc && v, true)
}