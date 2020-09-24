def listsum(numList):
    if len(numList) == 1:
        return numList[0]
    else:
        return numList[0] + listsum(numList[1:])


print(listsum([1, 3, 5, 7, 9]))

"""
上面程序要点：
1.问题分解为更小规模的相同问题，并表现为调用自身
2.对最小规模的问题的解决： 简单直接

递归3定律
 递归算法必须有个基本结束条件
 递归算法必须改变自己的状态并向基本结束条件演进
 递归算法必须递归地调用自身
"""


"""
递归深度
sys.setrecursionlimit(num)
"""