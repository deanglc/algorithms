package main

import "fmt"

// 今天才发现平时使用的是slice而不是list，整理一下slice，list和映射
// 函数间传递数组开销很大，因为函数属于值传递，大数组尽量以指针传递
// 1.数组长度不可变， [...]int{1,2,3}length=3 使用...替代数组的长度，
// Go 语言会根据初始化时数组元素的数量来确定该数组的长度
/*
var l1 [5]int{1,2,3,4,5}
var l [5]int
l3 := [5]int{3:666} //[0,0,0,666,0]
array := [...]int{10, 20, 30, 40, 50}  //以上声明的是数组，长度不可改变

slice := []int{1,2,3} //此为slice，长度、容量均可变，append方法
*/
// golang多维数组声明
// var duoList [4][2]int  [[0,0] [0,0] [0,0] [0,0]]
// array := [4][2]int{{10, 11}, {20, 21}, {30, 31}, {40, 41}}
// array := [4][2]int{1: {20, 21}, 3: {40, 41}}
// array[1][0] 值20

func main() {
	list := [5]int{1: 10, 0: 10}

	fmt.Print(list)

	l := [5]*int{0: new(int), 1: new(int)}
	fmt.Println(l)
	fmt.Println("\n", *l[0], &l[0], &l[1])

	*l[0] = 1
	fmt.Println(*l[0], &l[0])

	ll := l
	fmt.Println(ll, l)
	*ll[1] = 2
	fmt.Println(ll, l)
	fmt.Println(*ll[1], *l[1])

	a := [...]int{1, 2, 3}
	b := a
	fmt.Println(a, b)
	b[0] = 666
	fmt.Println(a, b)
	fmt.Printf("%p %p", &a, &b)
	ad := []int{1, 2, 3, 4, 5}
	adc := ad[1:2:5]
	fmt.Println("\n", adc, len(adc), cap(adc))
	fmt.Println("\n", adc[3])
}
