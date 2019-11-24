package main

import (
    "fmt"
    "time"
)

func main() {
    timer := time.NewTimer(time.Second)

    <- timer.C
    fmt.Println(time.Now())
}