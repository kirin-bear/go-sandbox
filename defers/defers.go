package defers

import "fmt"

var Global = 5 

func Start() {
	fmt.Println("Работа с defer")
	task1()
	//fmt.Println("Global after task1", Global)
}

// Директиву defer обычно применяют для освобождения ресурсов — например, чтобы закрыть каналы, файлы, сетевые соединения после окончания работы. Об использовании системных ресурсов расскажем в следующих темах курса, а пока рассмотрим в качестве ресурса глобальную переменную, доступную всем функциям программы:
// Используя defer, напишите функцию, которая меняет эту переменную и выводит на экран её новое значение. Потом она должна вернуть всё как было.
// 
// 
func task1() {
	//fmt.Println("Работа task1")
	//fmt.Println("Global before defer", Global)
	//newGlobal := 4
	//defer func (v int)  {
	//	Global = v
	//	fmt.Println("Global after defer", Global)
	//}(newGlobal)
	fmt.Println(Global)
    UseGlobal()
    fmt.Println(Global)
	
}

func UseGlobal() {
    defer func(checkout int) {
        Global = checkout // присваиваем Global значение аргумента
    }(Global) // копируем значение Global в аргументы функции
    Global = 42 // Изменяем Global  
    fmt.Println(Global)
    // Здесь будет вызвана отложенная функция
}