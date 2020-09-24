package main

import (
	"fmt"
	"time"
)

func grt1(c chan int32) {
	for i := 'a'; i <= 'z'; i++ {
		c <- i
		//fmt.Println("g-1:", i)
		if i%2 == 1 {
			fmt.Println("g-1:", i)
		}
	}
}

func grt2(c chan int32) {
	for i := 'a'; i <= 'z'; i++ {
		<-c
		//fmt.Println("g-2:", i)

		if i%2 == 0 {
			fmt.Println("g-2:", i)
		}
	}
}

func main() {
	t1 := time.Now().UnixNano()
	msg := make(chan int32)
	go grt1(msg)
	go grt2(msg)
	time.Sleep(1 * time.Second)
	fmt.Println(time.Now().UnixNano() - t1)
}
