#!/usr/bin/env python3
"""
 Author: Kevin Turkington
 2/8/18
 References:
 [0] ProcessNumber https://en.wikipedia.org/wiki/AKS_primality_test
 [1] getDigitSum https://stackoverflow.com/questions/14939953/sum-the-digits-of-a-number-python
 Description:
"""
from math import sqrt

class solution():
    # range 1..100001 creates inclusive set.
    def __init__(self,r=range(1,100001)):
        self.r = r

        lastElement = self.r[-1] -1 # if range is default grab 99999
        lastElementSum = self.getDigitSum(lastElement) +1 # grab max value for digits
        self.memorization = [None]*lastElementSum # fixed size array initialized with -1

    def ProcessNumber(self, number):
        """
        Checks if number is prime and stores prime number in a memorization list.
        Ref: [0]
        """
        # if even stop cannot be prime or if number has already been calculated.
        if number % 2 == 0 or self.memorization[number] != None:
            return
        
        rangelimit = int( sqrt(number) +1 )
        for i in range(3, rangelimit, 1):
            if number % i == 0:
                return
        self.memorization[number] = number

    def getDigitSum(self, number):
        """
        Gets the sum of digits in a number.
        Ref: [1]
        """
        total = 0

        while number > 0:
            total += number % 10 # add a digit
            number //= 10 # divide by tenth, hundredth... and grab floor of division
        return total
    
    def run(self):
        """
        Sums the digits of each item in the list
        Checks if the sum is a prime number
        if items position in memorization map is not none print its value.
        """
        for i in range(1, len(self.r)):
            itemSum = self.getDigitSum(self.r[i]) # generate sum
            self.ProcessNumber(itemSum)

            if self.memorization[itemSum] != None:
                print(self.memorization[itemSum])

if __name__ == "__main__":
    solution().run()