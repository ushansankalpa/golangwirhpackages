package main

import (
	"encoding/json"
	"fmt"
)

//first
type Book struct {
	Title  string `json:"title"`
	Author Author `json:"author"`
}

type Author struct {
	Name      string `json:"name"`
	Age       int    `json:"age"`
	Developer bool   `json:"is_developer"`
}

func main() {

	author := Author{Name: "Ushan", Age: 21, Developer: true}

	book := Book{Title: "Learnig go", Author: author}
	fmt.Printf("%+v\n", book)

	byteArray, err := json.MarshalIndent(book, "", " ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(byteArray))

}
