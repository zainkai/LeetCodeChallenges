/**
 * Initialize your data structure here.
 */
var Logger = function() {
  this.map = {}
};

/**
* Returns true if the message should be printed in the given timestamp, otherwise returns false.
      If this method returns false, the message will not be printed.
      The timestamp is in seconds granularity. 
* @param {number} timestamp 
* @param {string} message
* @return {boolean}
*/
Logger.prototype.shouldPrintMessage = function(timestamp, message) {
  for(let i = timestamp - 9; i <= timestamp; i++) {
      if (this.map[i] && this.map[i].indexOf(message) !== -1) {
          return false
      }
  }
  
  if(this.map[timestamp]) {
      this.map[timestamp].push(message)
  } else {
      this.map[timestamp] = [message]
  }
  
  return true
};

/** 
* Your Logger object will be instantiated and called as such:
* var obj = new Logger()
* var param_1 = obj.shouldPrintMessage(timestamp,message)
*/