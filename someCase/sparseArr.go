package main

import (
	"fmt"
	"unsafe"
)

// refer: https://www.cnblogs.com/louis181214/p/10412889.html

// 稀疏数组-sparseArr 使用场景
// 1.当一个数组(n维)大部分元素为0或者默认值
// 2.把不同的元素的行列值记录在一个小规模数组中，从而缩小程序的规模

// 以下为五子棋 棋盘可以用稀疏数组实现

/*
0 0 1 0 0 0 0 0 0 0 0
0 1 0 2 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0 0 0
1:黑 2:白
以上共三个棋子
黑子坐标 (0,2) (1,1)  白子坐标(1,3)
*/

type ValNode struct {
	row int //行
	col int //列
	val int
}

func main() {
	// 对比 数组实现棋盘
	arrMap := [11][11]int{}

	arrMap[0][2] = 1
	arrMap[1][1] = 1
	arrMap[1][3] = 2

	for _, v1 := range arrMap {
		for _, v2 := range v1 {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println() //换行
	}

	//稀疏数组 保存棋盘
	var sparseArr []ValNode
	for row, v1 := range arrMap {
		for col, v2 := range v1 {
			if v2 != 0 {
				valNode := ValNode{
					row: row,
					col: col,
					val: v2,
				}
				sparseArr = append(sparseArr, valNode)
			}
		}
	}
	fmt.Println(sparseArr, unsafe.Sizeof(sparseArr), unsafe.Sizeof(arrMap))

	for i, vn := range sparseArr {
		fmt.Printf("%d: %d-%d-%d \n", i, vn.row, vn.col, vn.val)
	}
	//只需要保存棋盘大小 11-11-0 即可还原棋盘
}
