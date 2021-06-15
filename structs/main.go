package main

import (
	"fmt"
	"structs/computer"
)

func main() {
	spec := computer.Spec{
		Maker: "apple",
		Price: 200000,
		Model: "Mac min",
	}
	fmt.Println("Maker : ", spec.Maker)
	fmt.Println("Price : ", spec.Price)
}
