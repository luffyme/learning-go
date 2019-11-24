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
	group := ColorGroup{
		ID: 1,
		Name: "redis",
		Colors: []string{"crimson", "red", "ruby", "maroon"},
	}

	j, err := json.Marshal(group)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(string(j))
}