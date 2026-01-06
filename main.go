package main

import (
	"fmt"
	"strings"
)

func main() {
	var helloWorld, helloGo string
	var year, month, day int
	var birthday = 30

	helloWorld = "Hello, World"
	helloGo = "Hello, Go"
	year, month, day = 2026, 1, 5
	fmt.Println(helloWorld)
	fmt.Println(helloGo)
	fmt.Println("Current date:", year, "-", month, "-", day)
	fmt.Println("Birthday:", birthday)
	fmt.Println("==============")
	fmt.Println(strings.Split(helloGo, "")[0])
	fmt.Println(string(helloGo[0]))

	var i int
	var i32 int32
	fmt.Println(i == i32)
}
