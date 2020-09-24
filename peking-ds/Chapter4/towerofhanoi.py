def moveTower(height, fromPole, toPole, withPole):
    if height >= 1:
        moveTower(height - 1, fromPole, withPole, toPole)
        moveDisk(fromPole, toPole)
        moveTower(height - 1, withPole, toPole, fromPole)


def moveDisk(fp, tp):
    print("moving disk from", fp, "to", tp)


moveTower(3, "A", "B", "C")


# 上面的汉诺塔解法简直弟弟
def hanoi(n, a, b, c):
    if n == 1:
        print(a, '-->', c)
    else:
        hanoi(n - 1, a, c, b)
        print(a, '-->', c)
        hanoi(n - 1, b, a, c)


# 调用
hanoi(3, 'A', 'B', 'C')

"""
汉诺塔递归思路  a柱-开始柱 b-中间柱 c-目标柱
1. 将a柱子上的n-1个盘 从a经过c移到b   hanoi(n-1,a,c,b)
2. 将a柱子上的最大盘 从a移到c         a ---> c
3. 最后将b上的n-1个盘子 从b经过a移到c  hanoi(n-1,b,a,c)
"""