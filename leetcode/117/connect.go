package _17

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// v1 4ms-75.91 6.2mb-10.22
func connect2(root *Node) *Node {
	if root == nil {
		return nil
	}
	queue := []*Node{root}
	for len(queue) > 0 {
		tmp := queue
		queue = nil
		for i := 0; i < len(tmp); i++ {
			if i+1 < len(tmp) {
				tmp[i].Next = tmp[i+1]
			}
			// 添加下一层 循环往复 直到叶子节点层
			if tmp[i].Left != nil {
				queue = append(queue, tmp[i].Left)
			}
			if tmp[i].Right != nil {
				queue = append(queue, tmp[i].Right)
			}
		}
	}
	return root
}

// ================================================================================
func connect(root *Node) *Node {
	n := root
	var prev, nextLevel *Node // prev是被连接这层的上一次成功连接的节点，nextLevel是这层的第一个节点。
	for n != nil {            // 来迭代吧，一直到节点世界的尽头。。。。。
		prev, nextLevel = nil, nil // 下一层目前还没有成功连接Next的节点，也找到下一层的第一个节点
		for n != nil {             // 循环到本层没节点为止
			connectNode(n.Left, &prev, &nextLevel)  // 尝试连接本层当前节点的Left child节点
			connectNode(n.Right, &prev, &nextLevel) // 尝试连接本层当前节点的Right child节点
			n = n.Next                              // 当前节点的children节点处理完，处理下一个节点
		}
		n = nextLevel // 这层的children都连接完了，那么该通过相同套路迭代刚刚连接完Next的这层了
	}
	return root
}

func connectNode(n *Node, prev, nextLevel **Node) {
	switch {
	case n == nil: // 这个节点是nil，直接返回
		return
	case *nextLevel == nil:
		*nextLevel = n // 这个节点是这层第一个不是nil的节点，也就是首节点
		*prev = n      // 因此也是第一个我们见到的活着的非nil节点
		return
	}
	(*prev).Next = n // 之前这层已经连接到了prev节点，那么prev节点的Next指向这个节点吧
	*prev = n        // 更新已经被连接完的这个节点为prev节点
}
