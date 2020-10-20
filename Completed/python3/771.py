class Solution:
    def numJewelsInStones(self, J: str, S: str) -> int:
        jewels = {}
        for j in J:
            jewels[j] = True
            
        result = 0
        for s in S:
            if s in jewels:
                result+=1
        return result
        
