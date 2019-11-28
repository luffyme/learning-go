package main

import (
	"fmt"
	"strings"
	"net/url"
)

func main() {
	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	u, err := url.Parse(s)
    if err != nil {
        fmt.Println("url.Parse error : ", err)
	}
	
	fmt.Println(u.Path)
    fmt.Println(u.Fragment)
	fmt.Println(u.Scheme)
	fmt.Println(u.User)
	fmt.Println(u.User.Username())

	fmt.Println(u.User.Password())
	pass, _ := u.User.Password()
	fmt.Println(pass)

	fmt.Println(u.Host)
    h := strings.Split(u.Host, ":")
    fmt.Println(h[0])
	fmt.Println(h[1])
	
	fmt.Println(u.RawQuery)
    m, _ := url.ParseQuery(u.RawQuery)
    fmt.Println(m)
    fmt.Println(m["k"][0])
}