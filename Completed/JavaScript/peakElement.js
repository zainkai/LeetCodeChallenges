/**
 * find the index of the element that is larger than its neighbors
 * nums[-1] = nums[nums.length] = Negative Infinity
 * @param {int} num 
 */
function findPeak(num) {
  if(num.length === 1) return 0
  if(num.length === 2) return num[0] > num[1] ? 0 : 1
  let low = 0
  let high = num.length -1
  
  while(low <= high) {
    let mid = Math.floor((low + high) /2)
    if (num[mid] > num[mid +1] && num[mid] > num[mid -1]) return mid
    if (mid === num.length -1 && num[mid] > num[mid -1]) return mid
    if (mid === 0 && num[mid] > num[mid +1]) return mid

    if (num[mid] > num[mid+1]) {
      high = mid -1
    } else {
      low = mid +1
    }
  }
  return 0
}


const arr = [1,2,1,3,5,6,4]

console.log(findPeak(arr))