package main

import (
	"fmt"
)

func change(s ...int) {
	s = append(s,3)
	fmt.Println(s)
}

func main() {
    slice := make([]int,5,5)
    slice[0] = 1
	slice[1] = 2
	fmt.Println(slice)
	// {1, 2, 0, 0, 0}
	change(slice...)				//slice还是原来的slice。在change中，因为扩容，新建了一个slice也就是s
	fmt.Println(slice)
	//{1, 2, 0, 0, 0}
	change(slice[0:2]...)
	//{1, 2}
	fmt.Println(slice)
	//{1, 2, 3, 0, 0}


	/////////////////

    var a = []int{1, 2, 3, 4, 5}
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
}