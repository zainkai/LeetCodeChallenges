function kEmptySlot(flowers, k) {
  let positions = {}
  for (let i = 0; i < flowers.length; i++) {
    positions[flowers[i] -1] = i + 1
  }

  console.log(positions)

  for (let i = 0; i < flowers.length; i++) {
    if (positions[i] - positions[i + 1] === k) {
      return i + 1
    }
  }
  return -1
}