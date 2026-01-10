package fizzbuzz

import (
	"fmt"
	"strconv"
)

func Start() {

	fmt.Println("FizzBuzz Start");

	for i:=1; i<=100; i++ {
		str := ""

		if (i%3 == 0) {
			str += "Fizz"
		}

		if (i%5 == 0) {
			str += "Buzz"
		}

		
		str = strconv.Itoa(i) + ". " + str
		fmt.Println(str);
	}
}