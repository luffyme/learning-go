package main

import (
    "crypto/md5"
    "crypto/sha1"
    "crypto/sha256"
    "io"
    "fmt"
    "encoding/hex"
)

func main() {
    md5Str := md5Func("123456")
    fmt.Println(md5Str)

    sha1Str := sha1Func("123456")
    fmt.Println(sha1Str)

    sha256Str := sha256Func("123456")
    fmt.Println(sha256Str)
}

func md5Func(str string) string {
    h := md5.New()
    io.WriteString(h, str)
	cipherStr := h.Sum(nil)
    return hex.EncodeToString(cipherStr)
}

func sha1Func(str string) string {
    h := sha1.New()
    io.WriteString(h, str)
    cipherStr := h.Sum(nil)
    return hex.EncodeToString(cipherStr)
}

func sha256Func(str string) string {
    h := sha256.New()
    io.WriteString(h, str)
    cipherStr := h.Sum(nil)
    return hex.EncodeToString(cipherStr)
}