class Solution:
    res = []
    def fizzBuzz(self, n: int) -> List[str]:
        self.res =[]
        for i in range(1,n+1):
            s = ""
            if i % 3 == 0:
                s += "Fizz"
            if i % 5 == 0:
                s += "Buzz"
            if i % 3 != 0 and i % 5 != 0:
                s += str(i)
                
            self.res.append(s)
        
        return self.res
        
        
