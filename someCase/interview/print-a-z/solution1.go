/* https://github.com/stong1994/interview/tree/master/Golang-2019-2-20
打印出a,b,c,d.....z
goroutine1: print(a,c,e,....)
goroutine2: print(b,d,f,....)
*/
package main

import (
	"fmt"
)

// 控制主线程+控制顺序输出
func main() {
	boolChan := make(chan bool)
	endChan := make(chan struct{})
	defer close(endChan)
	defer close(boolChan)

	go func() {
		for i := 'a'; i <= 'z'; i += 2 {
			if <-boolChan {
				fmt.Println("ace - ", string(i))
			}
			boolChan <- false
		}
	}()

	go func() {
		for j := 'b'; j <= 'z'; j += 2 {
			if !<-boolChan {
				fmt.Println("bdf - ", string(j))
			} else { //如果先运行这个goroutine 让j回滚到起点即可防止出现 a,d,c,f。。。的乱序情况
				j -= 2
			}
			if j == 'z' {
				endChan <- struct{}{}
			}
			boolChan <- true
		}
	}()

	boolChan <- true
	<-endChan
	//<-endChan

	//a := struct {}{}
	//b := true
	//fmt.Println(unsafe.Sizeof(a),unsafe.Sizeof(b))
}
