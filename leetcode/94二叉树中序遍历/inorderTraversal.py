class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


def inorderTraversal(self, root: TreeNode) -> list[int]:
    stack, rst = [root], []
    while stack:
        i = stack.pop()
        if isinstance(i, TreeNode):
            stack.extend([i.right, i.val, i.left])
        elif isinstance(i, int):
            rst.append(i)
    return rst
