package main

import (
	"fmt"
	"sync"
)

func main() {
	var mutex = &sync.RWMutex{}

	//加读锁，其他也可以获取读锁，但是不能获取写锁
	mutex.RLock()
	fmt.Println("RLock")
	mutex.RUnlock()

	//加读写锁，其他不能获取读锁或者读写锁
	mutex.Lock()
	fmt.Println("Lock")
	mutex.Unlock()
}