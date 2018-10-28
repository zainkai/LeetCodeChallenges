function countRecurringChar(str) {
  const map = {};
  (str).split("").forEach(e => {
    if (map[e] == null) {
      map[e] = 1
    } else {
      map[e] += 1
    }
  })

  return Object.keys(map).toString()
}

function recurringChar(str) {
  return Array.from(new Set((str).split("")))
}