function fizzBuzz (n, result = []) {
  if (n === 0) return result.reverse()
  if(n % 3 === 0 && n % 5 === 0){
    result.push("FizzBuzz")
    return fizzBuzz(n-1, result)
  }
  if(n % 3 === 0) {
    result.push("Fizz")
    return fizzBuzz(n-1, result)
  }
  if (n % 5 === 0) {
    result.push("Buzz")
    return fizzBuzz(n-1, result)
  }

  result.push(String(n))
  return fizzBuzz(n-1, result)
}

console.log(fizzBuzz(15))