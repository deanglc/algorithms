"""
给定一组不含重复元素的整数数组nums，返回该数组所有可能的子集（幂集）。
说明：解集不能包含重复的子集。

示例:
输入: nums = [1,2,3]
输出:
[
  [3],
  [1],
  [2],
  [1,2,3],
  [1,3],
  [2,3],
  [1,2],
  []
]
"""


class Solution(object):
    @staticmethod
    def subsets(nums: [int]) -> [[int]]:
        output = [[]]

        for num in nums:
            print(output)
            output += [curr + [num] for curr in output]

        return output


class Solution2(object):
    @staticmethod
    def subsets(nums: [int]) -> [[int]]:
        def back_track(first=0, curr=[]):
            if len(curr) == k:
                output.append(curr[:])
            for i in range(first, n):
                curr.append(nums[i])
                back_track(i + 1, curr)
                curr.pop()

        output = []
        n = len(nums)
        for k in range(n + 1):
            back_track()
        return output


if __name__ == '__main__':
    a = Solution2()
    print(a.subsets([1, 2, 3]))
