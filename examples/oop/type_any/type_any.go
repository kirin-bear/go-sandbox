package main

import (
	"fmt"
	"strings"
)

// Попробуйте реализовать обобщение операции умножения для чисел и строк. Если первый аргумент функции — строка, то повторить её b раз, а если число, то вернуть a*b.
func Mul(a any, b int) any {
	
	switch a2 := a.(type) {
	case int:
		return a2 * b 
	case string:
        return strings.Repeat(a2, b)
    case fmt.Stringer:
        return strings.Repeat(a2.String(), b)
    default :
        return nil
	}
}

func main() {
	fmt.Println(Mul(2,2));
	fmt.Println(Mul("2",4));
}