function validParenthesis(str) {
  const parenMap = {
    '(': ')',
    '{': '}',
    '[': ']'
  }
  const stk = []
  str.split("").forEach(paren => {
    const topStk = stk[stk.length -1]
    if (stk.length === 0 || parenMap[topStk] !== paren) {
      stk.push(paren)
    } else {
      stk.pop()
    }
  })

  return !stk.length
}

const result = validParenthesis("()[]{}")
console.log(result)