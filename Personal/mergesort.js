function mergeSort(unSortedArr) {
  if (unSortedArr.length === 1) { return unSortedArr }

  const midIdx = Math.floor(unSortedArr.length /2)
  const leftHalf = unSortedArr.slice(0, midIdx)
  const rightHalf = unSortedArr.slice(midIdx)

  return mergeHelper(
    mergeSort(leftHalf),
    mergeSort(rightHalf)
  )
}

function mergeHelper(leftArr, rightArr){
  let leftIter = 0
  let rightIter = 0
  let result = []
  while(leftIter < leftArr.length && rightIter < rightArr.length) {
    if(leftArr[leftIter] < rightArr[rightIter]) {
      result.push(leftArr[leftIter])
      leftIter++
    } else {
      result.push(rightArr[rightIter])
      rightIter++
    }
    console.log(result)
  }
  result = result.concat(leftArr.slice(leftIter))
  result = result.concat(rightArr.slice(rightIter))
  return result
}

console.log(mergeSort([5,7,9,3,0,1]))