package main

import (
	"fmt"
	"time"
)

func tel(c chan int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
}

func main() {
	var chInt chan int = make(chan int)

	go tel(chInt)

	for {

		time.Sleep(time.Second)

		select {
		case o, ok := <-chInt:
			if !ok {
				fmt.Println("over")
				break
			}
			fmt.Println(o)
		default:
			fmt.Println("default")
		}

	}
}
