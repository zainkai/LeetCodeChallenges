// this exact same code works in my es6 enviroment but not leetcode. leetcode == mathlab?

function mergeIntervals(intervals) {
  intervals.sort((a,b) => a[0] > b[0])
  for(let i = 0; i < intervals.length -1; i++) {
    // const [s1,e1] = intervals[i]
    // const [s2,e2] = intervals[i+1]

    const s1 = intervals[i][0], e1 = intervals[i][1]
    const s2 = intervals[i+1][0], e2 = intervals[i+1][1]

    if (e1 > s2) {
      intervals.splice(i,1)
      const i2 = intervals.findIndex(x => x[0]=== s2 && x[1] === e2)
      intervals.splice(i2,1)
      intervals.push([s1,e2])
      intervals.sort((a,b) => a[0] > b[0])
    }
  }
  return intervals
}

console.log(mergeIntervals([[1,3],[2,6],[8,10],[15,18]]))
