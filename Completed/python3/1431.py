class Solution:
    def kidsWithCandies(self, candies: List[int], extraCandies: int) -> List[bool]:
        maxCandy = self.getMax(candies)
        
        results = []
        for candy in candies:
            results.append(candy+extraCandies >= maxCandy)
        return results
            
    def getMax(self, arr):
        res = arr[0]
        for elm in arr:
            if elm > res:
                res = elm
        return res
    
        
