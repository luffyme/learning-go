package main

import (
    "fmt"
    "time"
)

func main() {
	//单位
	//time.Millisecond
	//time.Second

	//当前时间戳
	nowTime := time.Now()
	fmt.Println(nowTime.Unix())
	fmt.Println(nowTime.UnixNano())
	fmt.Println(nowTime.Format("2006-01-02 15:04:05")

	//各种时间
	now := time.Now()
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	fmt.Println(now.Nanosecond())
	fmt.Println(now.Location())
	fmt.Println(now.Weekday())

	//睡眠
	time.Sleep(time.Second * 2)

	//时间差
	then := time.Now()
	diff := then.Sub(now)
	fmt.Println(diff)
	fmt.Println(diff.Hours())
	fmt.Println(diff.Minutes())
	fmt.Println(diff.Seconds())
	fmt.Println(diff.Nanoseconds())

	//增加时间
	fmt.Println(then.Add(3).Unix())
	fmt.Println(then.Add(-3).Unix())

	fmt.Println("Done")
}