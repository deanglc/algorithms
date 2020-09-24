class Solution:
    def combinationSum(self, candidates: [int], target: int) -> [[int]]:
        candidates = sorted(candidates)

        ans = []

        def find(s, use, remain):
            for i in range(s, len(candidates)):
                c = candidates[i]
                if c == remain:
                    ans.append(use + [c])
                if c < remain:
                    find(i, use + [c], remain - c)
                if c > remain:
                    return

        find(0, [], target)

        return ans


if __name__ == '__main__':
    a = Solution()
    print(a.combinationSum([2, 3, 6, 7], 7))
    print(a.combinationSum([7, 3, 2], 18))
