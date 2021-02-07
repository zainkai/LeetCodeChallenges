# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def rightSideView(self, root: TreeNode) -> List[int]:
      result = []
      self.helper(root, result)
      return result
    def helper(self, root: TreeNode, lst: List[int], level: int = 0):
      if root == None:
        return
      
      if level >= len(lst):
        lst.append(root.val)
      else:
        lst[level]=root.val
      
      self.helper(root.left, lst, level+1)
      self.helper(root.right, lst, level+1)
        
