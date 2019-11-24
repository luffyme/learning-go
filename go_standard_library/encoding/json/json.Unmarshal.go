package main

import (
	"fmt"
	"encoding/json"
)

type Animal struct {
	Name string
	Order string
}

func main() {
	var animals []Animal
	var jsonBlob = []byte(`[
		{"Name": "Platypus", "Order": "Monotremata"},
		{"Name": "Quoll", "Order": "Dasyuromorphia"}
	]`)

	err := json.Unmarshal(jsonBlob, &animals)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Printf("%+v", animals)
}