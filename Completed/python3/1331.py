class Solution:
    def arrayRankTransform(self, arr: List[int]) -> List[int]:
      setRank = {}
      for v in arr:
        setRank[v] = 0
        
      ranks = []
      for k in setRank:
        ranks.append(k)
      
      ranks.sort()
      for i, val in enumerate(ranks):
        setRank[val] = i+1
      
      for i, val in enumerate(arr):
        arr[i] = setRank[val]
      return arr
      
    
        
