from pythonds.basic.queue import Queue

"""一群弟弟传土豆，鼓声停下时手拿热土豆的弟弟出列，多次反复直到只剩下一个弟弟"""


def hot_potato(namelist, num):
    """num: 固定的传递次数
       namelist: 一群弟弟
    """
    didi_queue = Queue()
    # 小朋友全部入列
    for name in namelist:
        didi_queue.enqueue(name)

    while didi_queue.size() > 1:
        for i in range(num):
            didi_queue.enqueue(didi_queue.dequeue())  # 队列滚动num次后

        didi_queue.dequeue()  # 队首的弟弟出列
    # 队列此时只剩下1个弟弟，出列即可
    return didi_queue.dequeue()


print(hot_potato(["Bill", "David", "Susan", "Jane", "Kent", "Brad"], 7))
