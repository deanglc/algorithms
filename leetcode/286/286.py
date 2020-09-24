"""
你被给定一个 m × n 的二维网格，网格中有以下三种可能的初始化值：
-1 表示墙或是障碍物
0 表示一扇门
INF 无限表示一个空的房间。然后，我们用 2**31 - 1 = 2147483647 代表 INF。你可以认为通往门的距离总是于 2147483647 的。
你要给每个空房间位上填上该房间到 最近 门的距离，如果无法到达门，则填 INF 即可。
示例：
给定二维网格：

INF  -1  0  INF
INF INF INF  -1
INF  -1 INF  -1
  0  -1 INF INF
运行完你的函数后，该网格应该变成：

  3  -1   0   1
  2   2   1  -1
  1  -1   2  -1
  0  -1   3   4
"""


class Solution:
    def wallsAndGates(self, rooms: list[list[int]]) -> None:
        """
        Do not return anything, modify rooms in-place instead.
        """
        if not rooms or len(rooms) == 0:
            return
        row = len(rooms)
        col = len(rooms[0])

        def bfs(rooms, i, j, val):  # val表示环扩展的距离
            nonlocal row, col
            if i < 0 or i >= row or j < 0 or j >= col:
                return
            if rooms[i][j] < val:  # 如果当前位置已经找到更近的门了，直接返回
                return
            rooms[i][j] = val
            # 广度 递归
            bfs(rooms, i + 1, j, val + 1)
            bfs(rooms, i, j + 1, val + 1)
            bfs(rooms, i - 1, j, val + 1)
            bfs(rooms, i, j - 1, val + 1)

        for i in range(row):
            for j in range(col):
                if rooms[i][j] == 0:
                    bfs(rooms, i, j, 0)


"""
317.离建筑物最近的距离 Hard
489.扫地机器人 Hard
994.腐烂的橘子 Easy
"""