# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def minDepth(self, root: TreeNode) -> int:
      if root == None:
        return 0
      return self.helper(root)
    def helper(self, root:TreeNode) -> int:
      if root.left == None and root.right == None:
        return 1
      
      val = 1<<63
      if root.left != None:
        val = min(val, self.helper(root.left)+1)
      if root.right != None:
        val = min(val, self.helper(root.right)+1)
        
      return val
