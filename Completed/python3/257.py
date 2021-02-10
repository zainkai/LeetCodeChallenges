# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def binaryTreePaths(self, root: TreeNode) -> List[str]:
      result = []
      if root == None:
        return result
      
      path = str(root.val)
      if root.left == None and root.right == None:
        return [path]
      
      self.dfs(root.left, path, result)
      self.dfs(root.right, path, result)
      
      return result
    
    def dfs(self, root: TreeNode, path: str, result: List[str]):
      if root == None:
        return
      
      path += "->" + str(root.val)
      if root.left == None and root.right == None:
        result.append(path)
        return
      
      self.dfs(root.left, path, result)
      self.dfs(root.right, path, result)
        
