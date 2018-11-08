class interval{
  constructor(s,e) {
    this.start = s
    this.end = e
  }
}

function minMeetingRooms(meetings) {
  if(meetings.length === 0) return 0
  if(meetings.length === 1) return 1
  
  const srtArr = meetings.map(e => e.start).sort()
  const endArr = meetings.map(e => e.end).sort()

  let maxRooms = 0
  let s = 0, e = 0, counter = 0
  while(s < srtArr.length) {
    if (srtArr[s] < endArr[e]) {
      s += 1
      counter += 1
    } else {
      e += 1
      counter -= 1
    }
    maxRooms = Math.max(maxRooms, counter)
  }
  return maxRooms
}

const arr = [new interval(7,10), new interval(2,4)]
console.log(minMeetingRooms(arr))