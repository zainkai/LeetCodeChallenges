/**
 * Initialize your data structure here.
 */
var HitCounter = function() {
  this.map = {}
};

/**
* Record a hit.
      @param timestamp - The current timestamp (in seconds granularity). 
* @param {number} timestamp
* @return {void}
*/
HitCounter.prototype.hit = function(timestamp) {
  if (this.map[timestamp]) {
      this.map[timestamp] += 1
  } else {
      this.map[timestamp] = 1
  }
};

/**
* Return the number of hits in the past 5 minutes.
      @param timestamp - The current timestamp (in seconds granularity). 
* @param {number} timestamp
* @return {number}
*/
HitCounter.prototype.getHits = function(timestamp) {
  let result = 0
  
  for (let key of Object.keys(this.map)) {
      const time = Number(key)
      if (timestamp - 300 < time && time <= timestamp) {
          result += this.map[key]
      }
  }
  
  return result
};

/** 
* Your HitCounter object will be instantiated and called as such:
* var obj = new HitCounter()
* obj.hit(timestamp)
* var param_2 = obj.getHits(timestamp)
*/