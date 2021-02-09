class Solution:
    def subtractProductAndSum(self, n: int) -> int:
        sumv, productv = 0, 1
        
        while n > 0:
            digit = n % 10
            n = int(n / 10)
            
            sumv += digit
            productv *= digit
        return productv - sumv
        
