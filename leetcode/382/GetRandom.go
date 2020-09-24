package main

// 类似398题
import (
	"fmt"
	"math/rand"
	"time"
)

/*
要求：
1.数据流长度N很大且不可知，所以不能一次性存入内存。
2.时间复杂度为O(N)。
3.随机选取m个数，每个数被选中的概率为m/N。
思路:
1. 接收的数据量小于m 则依次放入蓄水池
2. 当接收到第i个数据时，
	a.i>=m 在[0,i]范围内取随机数d
	b.若d落在[0,m-1]范围内，则用接收到的第i个数据 替换 蓄水池中的第d个数据
3.重复步骤2

算法的精妙之处在于：当处理完所有的数据时，蓄水池中的每个数据都是以m/N的概率获得的。
*/
func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	fmt.Println(r)
	s := rand.NewSource(time.Now().UnixNano()) // 使用当前的纳秒生成一个随机源，也就是随机种子
	ran := rand.New(s)
	fmt.Println(ran.Int()) // 获取随机数
	// 生成一个rand

}

type ListNode struct {
	Val  int
	Next *ListNode
}

type Solution struct {
	head *ListNode
	r    *rand.Rand
}

/** @param head The linked list's head.
  Note that the head is guaranteed to be not null, so it contains at least one node. */
func Constructor(head *ListNode) Solution {
	return Solution{
		head: head,
		r:    rand.New(rand.NewSource(time.Now().Unix())),
	}
}

/** Returns a random node's value. */
func (this *Solution) GetRandom() int {
	i := 1
	cur := this.head.Next
	val := this.head.Val
	for cur != nil {
		// 生成[0,i)之间的随机值， ----左闭右开
		if this.r.Intn(i) == 0 {
			// =0则替换蓄水池中的数据 因为
			val = cur.Val
		}
		i++
		cur = cur.Next
		// 循环往复 直到遍历完dataStream
	}
	return val
}

/*
m = 2  N=100
X = 2/i * (i+1-2/i+1 * i+2-2/i+2 * i+3-2/i+3 * i+4-2/i+4)
			i-1/i+1      i/i+2      i+1/i+3		i+2/i+4
			i-1/i+3 * i/i+4

*/
