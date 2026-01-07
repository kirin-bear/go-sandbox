package main

import "fmt"

func main() {

	/** 

	var height int
	var length int
	var weight float64
	var name   string
	var company = "Рога и копыта"

	// эквивалентно

	var (
		height, length int
		weight float64
		name   string
		company = "Рога и копыта"
	) 
	*/

	var (
		ver = "v0.0.1"
		id = 0
		pi = 3.1415
	)

	// Дополните пример объявлениями переменных ver, id, pi так, чтобы программа выводила ver = v0.0.1 id = 0 pi = 3.1415.
	fmt.Println("ver =", ver, "id =", id, "pi =", pi)
	


}