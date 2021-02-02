class Solution:
    def hammingWeight(self, n: int) -> int:
      bits = 0
      mask = 1
      while mask < (1<<63):
        if mask & n != 0:
          bits += 1
        mask <<= 1
      return bits
