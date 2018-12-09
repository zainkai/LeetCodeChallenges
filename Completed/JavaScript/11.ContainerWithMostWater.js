/**
 * @param {number[]} height
 * @return {number}
 */
var maxArea = function(height) {
  let max = -Infinity
  let L =0, R = height.length -1
  while(L < R) {
      const h = Math.min(height[L], height[R])
      const w = R - L
      max = Math.max(max, (h*w))
      if(height[L] < height[R]) {
          L +=1
      } else {
          R -=1
      }
  }
  return max
};