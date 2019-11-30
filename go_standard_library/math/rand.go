package main

import (
	"time"
	"fmt"
	"math/rand"
)

func main() {
	//rand.Intn 返回一个随机的整数 n，0 <= n <= 100。
	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Intn(100))

	//rand.Float64 返回一个64位浮点数 f，0.0 <= f <= 1.0。
	fmt.Println(rand.Float64())
	//这个技巧可以用来生成其他范围的随机浮点数，例如5.0 <= f <= 10.0
	fmt.Println((rand.Float64() * 5) + 5)
	fmt.Println((rand.Float64() * 5) + 5)

	//默认情况下，给定的种子是确定的，每次都会产生相同的随机数数字序列。要产生变化的序列，需要给定一个变化的种子。
	//变化的种子
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	fmt.Println(r1.Intn(100))

	//固定种子
	s2 := rand.NewSource(42)
	r2 := rand.New(s2)
	fmt.Println(r2.Intn(100))

	//如果种子相同，返回的随机也相同
	s3 := rand.NewSource(42)
	r3 := rand.New(s3)
	fmt.Println(r3.Intn(100))
}