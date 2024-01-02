package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var c chan int

func func1() {
	var i int
	defer wg.Done()
	timeout := time.After(20 * time.Second)
	for {
		select {
		case <-timeout:
			return

		default:
			if i == 5 {
				c <- 0
			}
			fmt.Printf("func1 - %v\n", i)
			i++
			time.Sleep(1 * time.Second)
		}
	}
}

func func2() {
	var i int
	defer wg.Done()
	timeout := time.After(20 * time.Second)
	for {
		select {
		case <-timeout:
			return

		default:
			if i == 5 {
				<-c
			}
			fmt.Printf("func2 - %v\n", i)
			i++
			time.Sleep(2 * time.Second)
		}
	}
}

func main() {
	wg.Add(2)
	c = make(chan int)
	go func1()
	go func2()
	wg.Wait()
}
