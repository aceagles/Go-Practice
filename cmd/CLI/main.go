package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	//defines the reader
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Simple Shell")
	fmt.Println("---------------------")

	for {
		fmt.Print("-> ")
		// This is where the input happens
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\r\n", "", -1)

		if strings.Compare("hi", text) == 0 {
			fmt.Println("hello, Yourself")
		}
		if v, err := strconv.Atoi(text); err == nil {
			fmt.Printf("Sleeping for %q seconds\n", text)
			// Async wait then come back
			go func(t int) {
				time.Sleep(time.Duration(t) * time.Second)
				fmt.Printf("Waking after %v seconds\n", t)
			}(v)
		}

	}

}
