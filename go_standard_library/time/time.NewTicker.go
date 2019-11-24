package main

import (
    "fmt"
    "time"
)

func main() {
    ticker := time.NewTicker(time.Second * 1)
    i := 0

    for {
        <- ticker.C //阻塞
        fmt.Println(time.Now())

        i++
        if i == 3 {
			//计数3后停止
            ticker.Stop() //停止定时器
            break //跳出循环
        }
	}
	
	/*
	for t := range ticker.C {
		fmt.Println(t)
	}
	*/

}