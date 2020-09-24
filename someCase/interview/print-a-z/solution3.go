package main

import (
	"fmt"
	"runtime"
)

// gorountine   a,b,c,d-----z   go  ace  go  bdf

func main() {
	runtime.GOMAXPROCS(12)
	aceChan := make(chan int32)
	bdfChan := make(chan int32)
	endChan := make(chan struct{}, 2)
	defer close(endChan)

	go ace(aceChan, bdfChan, endChan)
	go bdf(bdfChan, aceChan, endChan)

	aceChan <- 'a'

	<-endChan
	<-endChan
}

func ace(aceChan chan int32, bdfChan chan int32, endChan chan struct{}) {
	for {
		ch := <-aceChan
		if ch > 'y' {
			endChan <- struct{}{}
			close(aceChan)
			fmt.Println("aaa")
			ch += 1
			bdfChan <- ch
			break
		}
		fmt.Println("ace - ", string(ch))
		ch += 1
		bdfChan <- ch
	}
}

func bdf(bdfChan chan int32, aceChan chan int32, endChan chan struct{}) {
	for {
		ch := <-bdfChan
		if ch > 'z' {
			endChan <- struct{}{}
			close(bdfChan)
			fmt.Println("bbb")
			break
		}
		fmt.Println("bdf - ", string(ch))
		ch += 1
		aceChan <- ch
	}
}
