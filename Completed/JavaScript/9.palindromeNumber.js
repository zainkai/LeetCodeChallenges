function isPalinNumber(num) {
    if (typeof num !== 'number') return false
    const arr = String(num).split('')
    if (arr.length === 1) return true

    let l = 0, r = arr.length -1
    while(l < r) {
        if (arr[l] !== arr[r]) return false
        l++
        r--
    }
    return true
}

console.log(isPalinNumber(11))