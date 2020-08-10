class Solution:
    def isVowel(self, b:str):
        return b in ['a', 'e', 'i', 'o', 'u'] # O(5)
    def removeVowels(self, S: str) -> str:
        res = []
        for s in S:
            if not self.isVowel(s):
                res.append(s)
                
        return "".join(res)
