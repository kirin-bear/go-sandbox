package pointers

import "fmt"

func Start() {

	fmt.Println("Работа с указателями")

	i := 42
	p := &i
	fmt.Println(*p) // читаем значение переменной i через указатель 
	fmt.Println(p) // выведет пример формата - 0x140000100c8
	fmt.Println(i) // читаем значение переменной i через указатель 
	*p = 21         // записываем в переменную i значение 21 через указатель p
	fmt.Println(i) // выведет 21


	type A struct {
    	IntField int
	}

	p_a := &A{}
	p_a.IntField = 42 // вместо (*p).IntField = 42 
	fmt.Println(p_a) // выведет &{42}
	fmt.Println(*p_a) // выведет {42}

}