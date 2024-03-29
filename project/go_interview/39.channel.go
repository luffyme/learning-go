package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 100)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		} 
	}()

	go func() {
		for {
			a, ok := <-ch
			if !ok {
				fmt.Println("close")
				return
			}

			fmt.Println("a: ", a)
		} 
	}()

	close(ch)
	fmt.Println("ok")
	time.Sleep(time.Second * 10)
}