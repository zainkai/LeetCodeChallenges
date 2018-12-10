/**
 * @param {number[]} prices
 * @return {number}
 */
var maxProfit = function(prices) {
  let min = Infinity, profit = 0
  for(let i = 0;i < prices.length; i++) {
      if(min > prices[i]) min = prices[i]
      else {
          profit = Math.max(profit, prices[i] - min)
      }        
  }
  return profit
  
};