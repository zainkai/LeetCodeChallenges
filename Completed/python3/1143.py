class Solution:
    def longestCommonSubsequence(self, text1: str, text2: str) -> int:
        self.memo = {}
        self.t1 = text1
        self.t2 = text2
        
        tup = (len(text1), len(text2))
        return self.helper(tup)
    
    def helper(self, tup):
        
        if tup in self.memo:
            return self.memo[tup]
        elif tup[0] == 0 or tup[1] == 0:
            return 0
        elif self.t1[tup[0]-1] == self.t2[tup[1]-1]:
            x = (tup[0]-1, tup[1]-1)
            self.memo[tup] = 1+self.helper(x)
        else:
            a = (tup[0]-1, tup[1])
            b = (tup[0], tup[1]-1)
            
            self.memo[tup] = max(
                self.helper(a),
                self.helper(b))
            
        return self.memo[tup]
