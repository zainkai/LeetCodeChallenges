class Solution:
    def runningSum(self, nums: List[int]) -> List[int]:
        sum = 0
        res = []
        
        for n in nums:
            sum += n
            res.append(sum)
            
        return res
