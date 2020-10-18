class Solution:
    def defangIPaddr(self, address: str) -> str:
        result = []
        
        for ch in address:
            if ch == '.':
                result.append("[.]")
            else:
                result.append(ch)
                
        return "".join(result)
        
