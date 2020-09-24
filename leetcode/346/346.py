from collections import deque

"""给定一个整数数据流和一个窗口大小，根据该滑动窗口的大小，计算其所有整数的移动平均值。
示例
MovingAverage m = new MovingAverage(3);
m.next(1) = 1
m.next(10) = (1 + 10) / 2 
m.next(3) = (1 + 10 + 3) / 3
m.next(5) = (10 + 3 + 5) / 3 """


class MovingAverage(object):
    def __init__(self, size):
        self.deque = deque()
        self.max_size = size

    def next(self, value):
        self.deque.append(value)
        if len(self.deque) > self.max_size:
            self.deque.popleft()

        return sum(self.deque) / len(self.deque)


# class MovingAverage:
#
#     def __init__(self, size: int):
#         """
#         Initialize your data structure here.
#         """
#         from collections import deque
#         self.queue = deque(maxlen=size)
#
#     def next(self, val: int) -> float:
#         self.queue.appendleft(val)
#         return sum(self.queue) / len(self.queue)


if __name__ == '__main__':
    m = MovingAverage(3)
    print(m.next(1))
    print(m.next(10))
    print(m.next(5))
    print(m.next(3))
