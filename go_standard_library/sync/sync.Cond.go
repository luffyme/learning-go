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
    cond.Wait()   // 阻塞当前线程，直到收到该条件变量发来的通知, wati前必须先 lock
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
	cond.Signal() 					// 下发一个通知给已经获取锁的goroutine, 让该条件变量向至少一个正在等待它的通知的线程发送通知，表示共享数据的状态已经改变。
	
	time.Sleep(time.Second * 1)
    fmt.Println("Signal 2")
	cond.Signal() 					// 下发一个通知给已经获取锁的goroutine
	
	time.Sleep(time.Second * 1)
	fmt.Println("Broadcast")
	cond.Broadcast() 				//下发广播给所有等待的goroutine

	wg.Wait()
}
