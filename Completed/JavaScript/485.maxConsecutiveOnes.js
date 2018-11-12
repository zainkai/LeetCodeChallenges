function maxConsecOnes(nums) {
  if(nums.length === 0) return 0
  let max = 0
  
  for(let i = 0; i < nums.length; i++) {
    let count = 0
    for(let j  = i; j < nums.length; j++) {
      max = Math.max(max, count)
      flip = false
      if(nums[j] === 0 && flip) {
        break
      } else if(nums[j] === 0 && !flip) {
        flip = true
      }
      count += 1
    }
  }
  return max
}


console.log(maxConsecOnes([1,1,0,1,1,1]))