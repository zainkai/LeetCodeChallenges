class Solution:
    def largestUniqueNumber(self, A: List[int]) -> int:
        m = {}
        for v in A:
            if v in m:
                m[v]+=1
            else:
                m[v]=1
        
        res = -1
        for key, val in m.items():
            if val == 1 and key > res:
                res = key
        return res
