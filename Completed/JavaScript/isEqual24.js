function helper(numArr, acc) {
  const copy1 = [...numArr]
  if(numArr.length !== 0){
    copy1.splice(0,1)
    return helper(copy1,acc + numArr[0]) ||
    helper(copy1,acc - numArr[0]) ||
    helper(copy1,acc * numArr[0]) ||
    helper(copy1,acc / numArr[0])
  } else {
    console.log(acc)
    return acc === 24 ? true : false
  }
}

function numbersEqual24(fourNumbersArr) {
  let newArr = [...fourNumbersArr]
  return helper(newArr,0)
}

console.log(numbersEqual24([1,2,3,4]))