class BstNode {
  constructor(value, left, right){
    this.value = value
    this.left = left || null
    this.right = right || null
  }
}

function inOrderTraversal(root) {
  if (root === null) return

  inOrderTraversal(root.left)
  console.log(root.value)
  inOrderTraversal(root.right)
  return
}

function inOrderTraversalIter(root) {
  if (root === null) return
  let stack = []

  let curr = root
  while(curr !== null || stack.length > 0) {
    while(curr !== null) {
      stack.push(curr)
      curr = curr.left
    }
    curr = stack[stack.length -1]
    stack.pop()

    console.log(curr.value)
    curr = curr.right
  } 
}

function bstSearchRecur(root, target)

const root = new BstNode(300)
console.log(root.left)