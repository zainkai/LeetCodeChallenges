class Solution:
    def findMaxConsecutiveOnes(self, nums: List[int]) -> int:
      result = 0
      tmp = 0
      for v in nums:
        if v == 1:
          tmp+=1
        else:
          tmp=0
          
        result = max(result,tmp)
      return result
        
