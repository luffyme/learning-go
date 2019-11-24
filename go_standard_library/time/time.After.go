package main

import (
    "fmt"
    "time"
)

func main() {
	messages := make(chan string)

	select {
		case msg := <-messages:
			fmt.Println("received message", msg)
		//等待1秒超时
		case <-time.After(time.Second * 1):
			fmt.Println("timeout 1")
	}
}