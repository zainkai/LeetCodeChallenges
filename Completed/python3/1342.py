class Solution:
    def numberOfSteps (self, num: int) -> int:
        if num == 0:
            return 0
        elif num < 0:
            return 1<<63-1
        
        if num % 2 == 0:
            return self.numberOfSteps(num/2)+1
        else:
            return self.numberOfSteps(num-1)+1
        
