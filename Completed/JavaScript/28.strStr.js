/**
 * @param {string} haystack
 * @param {string} needle
 * @return {number}
 */
var strStr = function(haystack, needle) {
  if(haystack === needle) return 0
  for(let i = 0; i < haystack.length - needle.length +1; i++) {
      const slice = haystack.slice(i, i+needle.length)
      if (slice === needle) return i
  }
  return -1
};

/**
 * @param {string} haystack
 * @param {string} needle
 * @return {number}
 */
var strStr = function(haystack, needle) {
  return haystack.indexOf(needle)
};

