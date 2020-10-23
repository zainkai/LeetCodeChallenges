class Solution:
    def sortArrayByParity(self, A: List[int]) -> List[int]:
        evens = []
        odds = []
        
        for val in A:
            if val % 2 == 0:
                evens.append(val)
            else:
                odds.append(val)
        evens.extend(odds)
        return evens
        
