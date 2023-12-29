package main

import (
	"fmt"
	"time"
)

func state_machine(c, quit chan int) {
	for {
		select {
		case x := <-c:
			switch x {
			case 1:
				fmt.Println("Init")
				c <- 2
			case 2:
				fmt.Println("Connect")
				c <- 3
			case 3:
				fmt.Println("Check Connecction")

			}
		case <-quit:
			fmt.Println("Exiting")
			return
		}
	}
}

func main() {
	// This has to be buffered or the state machine can't send values to itself
	c := make(chan int, 2)
	quit := make(chan int, 1)

	go state_machine(c, quit)

	time.Sleep(5 * time.Second)

	c <- 1

	time.Sleep(5 * time.Second)
	c <- 1
	quit <- 0
	c <- 1
	time.Sleep(5 * time.Second)

}
