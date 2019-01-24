/**
 * @param {number[]} A
 * @return {number}
 */
var peakIndexInMountainArray = function(A) {
  return A.findIndex(x => x === Math.max.apply(null, A))
};

var peakIndexInMountainArray = function(A) {
  return A.findIndex(x => x === Math.max(...A))
};