package main

import "fmt"

func main() {
	var s string = "Hello 😆"
	fmt.Println(len(s))
	fmt.Println(s[6:10])
	fmt.Println(len("😆"))
}
