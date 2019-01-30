/**
 * // Definition for a Node.
 * function Node(val,children) {
 *    this.val = val;
 *    this.children = children;
 * };
 */
/**
 * @param {Node} root
 * @return {number[]}
 */
var preorder = function(root, arr = []) {
    if (root === null) return arr
    arr.push(root.val)
    root.children.forEach(c => {
        preorder(c, arr)
    })
    return arr
};
