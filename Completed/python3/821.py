class Solution:
    def shortestToChar(self, s: str, c: str) -> List[int]:
        result = [1<<62] * len(s)
        
        if s[0] == c:
            result[0] = 0
        for i in range(1, len(s), 1):
            ch = s[i]
            if ch == c:
                result[i] = 0
            else:
                result[i] = result[i-1]+1
        
        if s[len(s)-1] == c:
            result[len(s)-1] = 0
        for i in range(len(s)-2, -1, -1):
            ch = s[i]
            if ch == c:
                result[i] = 0
            else:
                result[i] = min(result[i], result[i+1]+1)
        return result
        
