# 找零问题  63分钱如何[1, 5, 10, 25]找零
# 1.递归法
def rec(coin_list, change):
    min_coins = change
    if change in coin_list:
        return 1
    else:
        for i in [c for c in coin_list if c <= change]:
            num_coins = 1 + rec(coin_list, change - i)
            if num_coins < min_coins:
                min_coins = num_coins

    return min_coins


# print(rec([1, 5, 10, 25], 63))


"""上面的方法 重复计算太多 进行了6000万次递归调用"""


def rec2(coin_list, change, known_results):
    min_coins = change
    if change in coin_list:
        known_results[change] = 1
        return 1
    elif known_results[change] > 0:
        return known_results[change]
    else:
        for i in [c for c in coin_list if c <= change]:  # [1 5 10]
            num_coins = 1 + rec2(coin_list, change - i, known_results)
            if num_coins < min_coins:
                min_coins = num_coins
                known_results[change] = min_coins

    return min_coins


"""修改后递归调用221次,因为缓存了部分结果"""


# print(rec2([1, 5, 10, 25], 63, [0] * 64))

# 2.动态规划 Dynamic Programming (DP)
def dp_make_change(coin_list, change, min_coins):
    for cents in range(1, change + 1):  # [1-11]
        coin_count = cents
        for j in [c for c in coin_list if c <= cents]:  # [1 5 10]
            if min_coins[cents - j] + 1 < coin_count:
                coin_count = min_coins[cents - j] + 1
        min_coins[cents] = coin_count
    return min_coins[change]


"""
DP主要思想：
1.从 最简单 的情况开始到达所需找零的循环
2.每一步都 依靠以前的最优解 来得到本步骤的最优解
3.直到得到答案
"""


# 动态规划找零最终版
def dp_make_change_last(coin_list, change, min_coins, coins_used):
    for cents in range(change + 1): # 计算1~change的各种最优解 假设change=63
        coin_count = cents  # 最差解=全部给 (1或者cents) 单位硬币
        new_coin = 1        # 给打印用
        for j in [c for c in coin_list if c <= cents]:   # 遍历找零货币列表
            if min_coins[cents - j] + 1 < coin_count:    #
                coin_count = min_coins[cents - j] + 1
                new_coin = j
        min_coins[cents] = coin_count
        coins_used[cents] = new_coin
    return min_coins[change]


def print_coins(coins_used, change):
    coin = change
    while coin > 0:
        this_coin = coins_used[coin]
        print(this_coin)
        coin = coin - this_coin


def main():
    amnt = 63
    clist = [1, 5, 10, 21, 25]
    coins_used = [0] * (amnt + 1)
    coin_count = [0] * (amnt + 1)

    print("Making change for", amnt, "requires")
    print(dp_make_change_last(clist, amnt, coin_count, coins_used), "coins")
    print("They are:")
    print_coins(coins_used, amnt)
    print("The used list is as follows:")
    print(coins_used)


if __name__ == '__main__':
    # print(dp_make_change([1, 5, 10, 21, 25], 63, [0] * 64))
    # print(dp_make_change([1, 5, 10, 25], 63, [0] * 64))
    # print(dp_make_change([1, 5, 10, 25], 28, [0] * 29))
    print(dp_make_change([1, 5, 10], 11, [0] * 12))
