class MinStack:

    def __init__(self):
        """
        initialize your data structure here.
        """
        self.stk = []
        self.min_stk = []

    def push(self, x: int) -> None:
        if len(self.stk) == 0:
          self.min_stk.append(x)
        else:
            self.min_stk.append(min(x, self.getMin()))
        self.stk.append(x)
        

    def pop(self) -> None:
      self.min_stk.pop()
      self.stk.pop()
        

    def top(self) -> int:
        if len(self.stk) == 0:
          return -1<<63
        return self.stk[len(self.stk)-1]

    def getMin(self) -> int:
        if len(self.min_stk) == 0:
          return -1<<63
        return self.min_stk[len(self.min_stk)-1]


# Your MinStack object will be instantiated and called as such:
# obj = MinStack()
# obj.push(x)
# obj.pop()
# param_3 = obj.top()
# param_4 = obj.getMin()
