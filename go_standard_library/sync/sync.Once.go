package main

import (
	"fmt"
	"sync"
)

func main() {
	var once sync.Once
	wg := sync.WaitGroup{}

	onceBody := func(){
		fmt.Println("only once")
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			once.Do(onceBody)
			defer wg.Done()
		}()
	}

	wg.Wait()
}