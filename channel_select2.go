package main

import (
	"fmt"
	"time"
)

func tel(c chan int, out chan int) {
	var old int = 0
	for {

		time.Sleep(time.Second)

		o, ok := <-c
		if !ok {
			break
		}
		if o > 100 {
			break
		}
		new := o + old

		c <- new
		out <- o

		old = o
	}
	close(c)
	close(out)
}

func main() {
	chIntt := make(chan int, 5)
	chResult := make(chan int)

	chIntt <- 1

	go tel(chIntt, chResult)

	for {
		time.Sleep(time.Second)
		select {
		case o, ok := <-chResult:
			if !ok {
				return
			}

			fmt.Println(o, ok)
		}

	}
}
