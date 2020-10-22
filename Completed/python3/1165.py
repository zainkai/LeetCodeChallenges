class Solution:
    def calculateTime(self, keyboard: str, word: str) -> int:
        keyDict = {}
        for i, key in enumerate(keyboard):
            keyDict[key]= i
            
        currIdx = 0
        time = 0
        for key in word:
            nextIdx = keyDict[key]
            time += abs(currIdx - nextIdx)
            currIdx = nextIdx
            
        return time
            
