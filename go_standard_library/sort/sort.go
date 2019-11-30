package main

import (
	"fmt"
	"sort"
)

type ByLength []string

func (s ByLength) Len() int {
    return len(s)
}
func (s ByLength) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}
func (s ByLength) Less(i, j int) bool {
    return len(s[i]) < len(s[j])
}

func main() {
	//对字符串切片进行排序
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	//对int切片进行排序
	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Ints:   ", ints)
	
	//判断是否已经排序好的
	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted: ", s)
	
	//通过自定义函数来排序
	fruits := []string{"peach", "banana", "kiwi"}
    sort.Sort(ByLength(fruits))
    fmt.Println(fruits)
}

