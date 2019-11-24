package main

import (
    "encoding/base64"
    "fmt"
)

// 编码
func base64Encode(str []byte) []byte {
    return []byte(base64.StdEncoding.EncodeToString(str))
}

// 解码
func base64Decode(str []byte) ([]byte, error){
    return base64.StdEncoding.DecodeString(string(str))
}

func main(){
    str := "hello"
    enc_str := base64Encode([]byte(str))
    fmt.Println(enc_str)
    fmt.Println(string(enc_str))

    dec_str,err := base64Decode(enc_str)
    if(err != nil){
        fmt.Println(err.Error())
    }

    fmt.Println(dec_str)
    fmt.Println(string(dec_str))
}