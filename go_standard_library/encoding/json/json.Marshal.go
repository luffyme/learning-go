package main

import (
	"fmt"
	"encoding/json"
)

type ColorGroup struct {
	ID int
	Name string
	Colors []string
}

func main() {


	//bool
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))
	
	//int
    intB, _ := json.Marshal(1)
	fmt.Println(string(intB))
	
	//float
    fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))
	
	//string
    strB, _ := json.Marshal("gopher")
    fmt.Println(string(strB))

	//slice
    slcD := []string{"apple", "peach", "pear"}
    slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))
	
	//map
    mapD := map[string]int{"apple": 5, "lettuce": 7}
    mapB, _ := json.Marshal(mapD)
    fmt.Println(string(mapB))

	group := ColorGroup{
		ID: 1,
		Name: "redis",
		Colors: []string{"crimson", "red", "ruby", "maroon"},
	}
	j, _ := json.Marshal(group)
	fmt.Println(string(j))
}