function successorBst(root, target) {
  if (root === null || target === null) return null
  
  let stk = inOrderHelper(root)
  let idx = stk.findIndex(e => e.value === target.value)
  if (idx === -1) return null
  return stk[idx + 1] || null
}

function inOrderHelper(root, stk=[]) {
  if (root === null) return

  inOrderHelper(root.left, stk)
  stk.push(root)
  inOrderHelper(root.right, stk)
  return stk
}
