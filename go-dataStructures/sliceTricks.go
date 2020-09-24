package main

import "fmt"

// https://github.com/EDDYCJY/blog
// https://juejin.im/post/5b95bf226fb9a05d16586851
// https://github.com/golang/go/wiki/SliceTricks
// https://github.com/go-community/gonote/blob/master/golang%E9%9D%A2%E8%AF%95%E9%A2%98/%E5%8D%B73.md
// https://studygolang.com/articles/18694
// https://yushuangqi.com/blog/2017/golang-mian-shi-ti-da-an-yujie-xi.html
// https://mp.weixin.qq.com/s/vn3KiV-ez79FmbZ36SX9lg

// golang.map在创建时准备足够空间有助于提升性能，
// 减少扩张时的内存分配和重新hash操作

// container/vector在go 1版本被移除
// Since the introduction of the append built-in, most of the functionality of the container/vector package, which was removed in Go 1, can be replicated using append and copy.
// Here are the vector methods and their slice-manipulation analogues:

var (
	s1 = []int{1, 2, 3}
	s2 = []int{11, 22, 33}
	sa = []string{"a", "b", "c"}
)

func main() {
	// func append(slice []Type, elems ...Type) []Type
	s1 = append(s1, s2...) // append单个元素，所以需要解包

	// Copy-1
	copySlice := make([]int, len(s1))
	fmt.Println(copySlice)

	copy(copySlice, s1) // 复制一份slice
	fmt.Println(copySlice, s1)
	fmt.Printf("copyS %p,s1 %p", copySlice, s1)

	s1[0] = 777 // copy复制了一份，并不是引用，so 结果如下
	fmt.Println("\n", copySlice, s1)
	// copy-2
	newS := append([]int(nil), s1...)
	fmt.Println(newS, s1)

	// copy-3
	newS3 := append(s1[:0:0], s1...)
	fmt.Println(newS3)

	fmt.Println(newS3[:0])
	fmt.Println(newS3[0:3:4]) // 这里的4为容量-cap，长度不可大于容量

	fmt.Println(len(newS3), cap(newS3))
	nn := newS3[:0:4]
	fmt.Println(len(nn), cap(nn))

	// cut
	s_cut := []int{1, 2, 3, 4, 5, 6}
	s_cut = append(s_cut[:2], s_cut[3:]...) // remove index=2
	fmt.Println(s_cut)

	// delete
	s_delete := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(s_delete[3:], s_delete[4:])
	copy(s_delete[3:], s_delete[4:])
	fmt.Println("delete==", s_delete)
}
