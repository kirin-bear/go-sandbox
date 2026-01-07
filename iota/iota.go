package main

import "fmt"

const (
    _ = iota*10  // обратите внимание, что можно пропускать константы 
    ten
    twenty
    thirty
)

const (
    hello = "Hello, world!"  // iota равна 0
    one = 1                  // iota равна 1

    black = iota   // iota равна 2
    gray
)

// задачка чтобы программа выводила 1 3 5 7 9 11.
// 
const (
	odin = iota*2 + 1
	tri
	pyat
	sem
	devyat
	odinnadcat
)

func main() {
    fmt.Println(ten, twenty, thirty) // 10, 20, 30
    fmt.Println(black, gray) // 2, 3
	fmt.Println("resolve example --------") // 2, 3
	fmt.Println(odin, tri, pyat, sem, devyat, odinnadcat)
} 