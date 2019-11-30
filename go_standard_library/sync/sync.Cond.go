package main

import (
	"fmt"
	"sync"
	"time"
)

var mutex = &sync.Mutex{}
var cond = sync.NewCond(mutex)
var wg = sync.WaitGroup{}

func test(x int) {
	defer wg.Done()

	fmt.Printf("start test %d\n", x)
	cond.L.Lock() // 获取锁
    cond.Wait()   // 等待通知 暂时阻塞
    fmt.Printf("start run %d\n", x)
    cond.L.Unlock()
}

func main() {
    for i := 0; i < 40; i++ {
		wg.Add(1)
        go test(i)
    }
	fmt.Println("start all")
	
	time.Sleep(time.Second * 1)
    fmt.Println("Signal 1")
	cond.Signal() 					// 下发一个通知给已经获取锁的goroutine
	
	time.Sleep(time.Second * 1)
    fmt.Println("Signal 2")
	cond.Signal() 					// 3秒之后 下发一个通知给已经获取锁的goroutine
	
	time.Sleep(time.Second * 1)
	fmt.Println("Broadcast")
	cond.Broadcast() 				//3秒之后 下发广播给所有等待的goroutine

	wg.Wait()
}
