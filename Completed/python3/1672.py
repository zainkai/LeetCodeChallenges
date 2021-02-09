import functools

class Solution:
    def maximumWealth(self, accounts: List[List[int]]) -> int:
        res = 0
        for banks in accounts:
            sumWealth = functools.reduce(lambda x,y : x+y, banks)
            res = max(res, sumWealth)
        return res
        
