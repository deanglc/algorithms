package main

import (
	"fmt"
)

func main() {
	c := make(chan bool, 1)
	c <- true

	end := make(chan struct{}, 2)

	go func() {
		for char := 'a'; char < 'a'+25; char++ {
			if <-c {
				fmt.Printf("%c", char)
				char++
			} else {
				char--
			}
			c <- false
		}
		end <- struct{}{}
	}()

	go func() {
		for char := 'b'; char <= 'b'+25; char++ {
			if !<-c {
				fmt.Printf("%c", char)
				char++
			} else {
				char--
			}
			c <- true
		}
		end <- struct{}{}
	}()

	<-end
	<-end
	fmt.Println("\n main end")

	//<-signal
	//time.Sleep(2*time.Second)
}
