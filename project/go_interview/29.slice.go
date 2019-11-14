package main

import (
	"fmt"
	"time"
)

func main() {
    v := []int{1, 2, 3}
    for i := range v {
        v = append(v, i)
	}
	fmt.Println(v)

	/////////////////////

    var m = [...]int{1, 2, 3}

    for i, v := range m {
        go func() {
            fmt.Println(i, v)
        }()
    }

    time.Sleep(time.Second * 3)
}