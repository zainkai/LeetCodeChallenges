class Solution:
    def maxPower(self, s: str) -> int:
      if len(s) <= 1:
        return len(s)
      
      tmp = 1
      result = 0
      
      for i in range(1, len(s)):
        prev = s[i-1]
        curr = s[i]
        
        if prev == curr:
          tmp+=1
        else:
          tmp = 1
        
        result = max(result, tmp)
      return result
