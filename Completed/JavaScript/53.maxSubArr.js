/**
 * @param {number[]} nums
 * @return {number}
 */
var maxSubArray = function(nums) {
  let max = -Infinity
  let curr = 0
  for(let i = 0; i < nums.length;i++) {
      curr = Math.max(curr + nums[i], nums[i])
      max = Math.max(max, curr)
  }
  return max
};