/**
 * // Definition for a Node.
 * function Node(val,children) {
 *    this.val = val;
 *    this.children = children;
 * };
 */
/**
 * @param {Node} root
 * @return {number}
 */
var maxDepth = function(root) {
  if (root === null) return 0
  let depth = 1
  for(let i = 0 ; i < root.children.length;i++) {
   let c = root.children[i]
   depth = Math.max(depth, 1 + maxDepth(c))
  }
  return depth
};