import random

"""P87-打印任务
平均每天任意一个小时有大约10个学生在实验室里,在这1小时中通常每人发起 2 次打印任务,每个打印任务的页数从 1 到 20 页不等。
草稿模式打印，每分钟打印10页；
高品质的打印模式，每分钟打印5页
较慢的打印速度可能会使学生等待太长时间。应该采取哪种打印模式?
我们关心的是每个学生等待打印的平均时长，这等于每个任务在队列里的等待平均时间。


"""
from pythonds.basic.queue import Queue


class Printer:
    def __init__(self, ppm):
        self.page_rate = ppm  # 打印速度
        self.currentTask = None
        self.timeRemaining = 0

    def tick(self):
        if self.currentTask != None:
            self.timeRemaining = self.timeRemaining - 1
            if self.timeRemaining <= 0:
                self.currentTask = None

    def busy(self):
        if self.currentTask != None:
            return True
        else:
            return False

    def startNext(self, new_task):
        self.currentTask = new_task
        self.timeRemaining = new_task.getPages() \
                             * 60 / self.page_rate


class Task:
    def __init__(self, time):
        self.timestamp = time  # 打印任务生成时间？
        self.pages = random.randrange(1, 21)

    def getStamp(self):
        return self.timestamp

    def getPages(self):
        return self.pages

    def waitTime(self, currenttime):
        return currenttime - self.timestamp


def simulation(numSeconds, pagesPerMinute):
    labprinter = Printer(pagesPerMinute)
    printQueue = Queue()
    waitingtimes = []

    for currentSecond in range(numSeconds):

        if newPrintTask():
            task = Task(currentSecond)
            printQueue.enqueue(task)

        if (not labprinter.busy()) and \
                (not printQueue.isEmpty()):
            nexttask = printQueue.dequeue()
            waitingtimes.append( \
                nexttask.waitTime(currentSecond))
            labprinter.startNext(nexttask)

        labprinter.tick()

    averageWait = sum(waitingtimes) / len(waitingtimes)
    print("Average Wait %6.2f secs %3d tasks remaining." \
          % (averageWait, printQueue.size()))


def newPrintTask():
    num = random.randrange(1, 181)
    if num == 180:
        return True
    else:
        return False


for i in range(10):
    simulation(3600, 5)
print("------------------------\n")
for i in range(10):
    simulation(3600, 10)
