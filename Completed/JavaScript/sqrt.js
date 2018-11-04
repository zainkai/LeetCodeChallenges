function sqrt(x) {
  let low = 0, high = Number.MAX_SAFE_INTEGER // not x
  let mid

  while (low <= high) {
    mid = Math.floor((low+high)/2)
    if ((mid * mid) === x) return mid
    if (mid > (x/mid)) {
      high = mid -1
    } else { 
      low = mid +1
    }
  }
  if((mid+1) > (x /(mid+1))) return mid
  return -1
}

console.log(sqrt(8))