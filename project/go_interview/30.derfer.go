package main

import (
	"fmt"
)

func f(n int) (r int) {
    defer func() {
        r += n
        recover()
    }()

    var f func()

    defer f()
    f = func() {
        r += 2
    }
    return n + 1
}

func main() {
	fmt.Println(f(3))
	
	////////////////

	var a = [5]int{1, 2, 3, 4, 5}
    var r [5]int

    for i, v := range a {
        if i == 0 {
            a[1] = 12
            a[2] = 13
        }
        r[i] = v
    }
    fmt.Println("r = ", r)
	fmt.Println("a = ", a)
	
	/*
		例子中参与循环的是 a 的副本，而不是真正的 a。就这个例子来说，假设 b 是 a 的副本，则 range 循环代码是这样的：
		for i, v := range b {
			if i == 0 {
				a[1] = 12
				a[2] = 13
			}
			r[i] = v
		}
	*/
}