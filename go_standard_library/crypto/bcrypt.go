package main

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func main() {
    for cost := 10; cost <= 15; cost++ {
        startedAt := time.Now()
		newPass, _ := bcrypt.GenerateFromPassword([]byte("password"), cost)
		
        duration := time.Since(startedAt)
        fmt.Printf("cost: %d, duration: %v, newPass: %s\n", cost, duration, string(newPass))
	}
	
	passwordOK := "admin"
	passwordERR := "adminxx"
	
    hash, err := bcrypt.GenerateFromPassword([]byte(passwordOK), bcrypt.DefaultCost)
    if err != nil {
        fmt.Println(err)
	}
	
	encodePW := string(hash)  // 保存在数据库的密码，虽然每次生成都不同，只需保存一份即可
    fmt.Println(hash)
    fmt.Println(encodePW)

    // 正确密码验证
    err = bcrypt.CompareHashAndPassword([]byte(encodePW), []byte(passwordOK))
    if err != nil {
        fmt.Println("pw wrong")
    } else {
        fmt.Println("pw ok")
    }

    // 错误密码验证
    err = bcrypt.CompareHashAndPassword([]byte(encodePW), []byte(passwordERR))
    if err != nil {
        fmt.Println("pw wrong")
    } else {
        fmt.Println("pw ok")
    }
}
