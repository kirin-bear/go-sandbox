package input_strings_counter

import (
    "bufio"
    "fmt"
    "os"
)

func Start() {
    // Получаем читателя пользовательского ввода
    reader := bufio.NewReader(os.Stdin)
	fmt.Println("Подсчет строк введенных пользователем")
    fmt.Println("Interaction counter")

    cnt := 0
    for {
        fmt.Print("-> ")
                // Считываем введённую пользователем строку. Программа ждёт, пока пользователь введёт строку 
        _, err := reader.ReadString('\n')
        if err != nil {
            panic(err)
        }

        f(&cnt)

        fmt.Printf("User input %d lines\n", cnt)
    }
}

func f(cnt *int) {
	*cnt += 1 // или *cnt++
}
