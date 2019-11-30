package main

import (
	"fmt"
	"bytes"
	"regexp"
)

func main() {
	//判断是否匹配，返回true或者false
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)
	
	r, _ := regexp.Compile("p([a-z]+)ch")
	//判断是否匹配，返回true或者false
	fmt.Println(r.MatchString("peach"))

	//返回匹配的字符串 返回string类型
	fmt.Println(r.FindString("peach punch"))

	//返回匹配的字符串的起始位置和结束位置 返回slice类型
	fmt.Println(r.FindStringIndex("peach punch"))

	//Submatch 返回完全匹配和局部匹配的字符串。返回slice类型
	fmt.Println(r.FindStringSubmatch("peach punch"))

	//返回完全匹配和局部匹配的索引位置。返回slice类型
	fmt.Println(r.FindStringSubmatchIndex("peach punch"))

	//返回所有的匹配项，而不仅仅是首次匹配项。例如查找匹配表达式的所有项。
	//如果第二个参数为正数，表示限制匹配的次数，-1表示不限制
	//返回slice类型
	fmt.Println(r.FindAllString("peach punch pinch", -1))
	fmt.Println(r.FindAllString("peach punch pinch", 1))

	//与上面一样，只不过是返回匹配的索引位置。返回slice类型
	fmt.Println(r.FindAllStringSubmatchIndex("peach punch pinch", -1))

	//参数为slice类型， 返回bool类型
	fmt.Println(r.Match([]byte("peach")))

	r = regexp.MustCompile("p([a-z]+)ch")
	//字符串替换
	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))
	
	//匹配的字符串进行函数处理，这里是转大写
	out := r.ReplaceAllFunc([]byte("a peach"), bytes.ToUpper)
    fmt.Println(string(out))
}

