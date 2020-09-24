"""
前中后序遍历区别： 经过root节点的顺序
"""
from pythonds.trees import binaryTree


def pre_order(tree: binaryTree):
    if tree:
        print(tree.getRootVal())
        pre_order(tree.getLeftChild())
        pre_order(tree.getRightChild())
