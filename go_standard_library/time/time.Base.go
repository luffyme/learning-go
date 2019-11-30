package main

import (
    "fmt"
    "time"
)

func main() {
	/*
		time.Nanosecond
		time.Microsecond
		time.Millisecond
		time.Second
		time.Minute
		time.Hour
	*/

	//当前时间戳
	nowTime := time.Now()
	fmt.Println("时间戳:")
	fmt.Println(nowTime.Unix())
	fmt.Println(nowTime.UnixNano())

	//格式化
	fmt.Println("格式化:")
	t,_ := time.ParseInLocation("2006-01-02 15:04:05", "2019-08-11 11:50:35", time.Local)
	fmt.Println(t.Unix())
	fmt.Println(t.Truncate(1 * time.Hour).Unix())
	fmt.Println(t.Truncate(1 * time.Hour)) 				// 整点（向下取整）
	fmt.Println(t.Round(1 * time.Hour))					// 整点（最接近）
	fmt.Println(t.Truncate(1 * time.Minute))			// 整分（向下取整）
	fmt.Println(t.Round(1 * time.Minute)) 				// 整分（最接近）


	//各种时间
	now := time.Now()
	fmt.Println("各种时间:")
	fmt.Println(now.String())
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	fmt.Println(now.Nanosecond())
	fmt.Println(now.Location())
	fmt.Println(now.Weekday())
	fmt.Println(now.Format("2006-01-02"))
	fmt.Println(now.Format("2006-01-02 15:04:05"))
	fmt.Println(now.Format(time.RFC3339)) 
	fmt.Println(now.Format("3:04PM"))
	fmt.Println(now.Format("Mon Jan _2 15:04:05 2006"))
	fmt.Println(now.Format("2006-01-02T15:04:05.999999-07:00"))
	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		now.Year(), now.Month(), now.Day(),
		now.Hour(), now.Minute(), now.Second())

	//年月日
	fmt.Println("年月日:")
	year, month, day := now.Date()
	fmt.Println(year, month, day)

	//睡眠
	fmt.Println("开始睡眠:")
	time.Sleep(time.Second * 2)

	//时间差
	then := time.Now()
	diff := then.Sub(now)
	fmt.Println("时间差值:")
	fmt.Println(diff)
	fmt.Println(diff.Hours())
	fmt.Println(diff.Minutes())
	fmt.Println(diff.Seconds())
	fmt.Println(diff.Nanoseconds())

	//增加时间
	fmt.Println("增加时间:")
	fmt.Println(then.Add(3).Unix())
	fmt.Println(then.Add(-3).Unix())

	fmt.Println("Done")
}