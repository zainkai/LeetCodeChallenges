/**
 * Definition for a binary tree node.
 * function TreeNode(val) {
 *     this.val = val;
 *     this.left = this.right = null;
 * }
 */
/**
 * @param {TreeNode} root
 * @return {boolean}
 */
var isUnivalTree = function(root) {
  return helper(root, root.val)
};

function helper(root, val) {
  if (root === null) return true
  return root.val === val && helper(root.left, val) && helper(root.right, val)
}