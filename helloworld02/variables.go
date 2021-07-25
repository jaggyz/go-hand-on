package main

import "fmt"

func main() {
	x := 10
	fmt.Println(x)
	k, v := 1, 1
	fmt.Println(k, v)
	var k1, v1 int = 1, 1
	fmt.Println(k1, v1)
	var (
		k2 int    = 1
		v2 string = "1"
	)
	fmt.Println(k2, v2)
}
