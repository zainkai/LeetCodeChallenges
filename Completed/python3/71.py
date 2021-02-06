class Solution:
    def simplifyPath(self, path: str) -> str:
      lst = path.split('/')
      
      result = []
      for tok in lst:
        if tok == '.' or tok == '/' or tok == '':
          continue
        elif tok == '..':
          if len(result) > 0:
            result.pop()
        else:
          result.append(tok)
      return '/' + "/".join(result)
      
        
