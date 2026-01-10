package functions

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"time"
)


func Start() {
	fmt.Println("Работа с функциями")


	countedPrint := countCall(myprint)
    countedPrint("Hello world")
    countedPrint("Hi")

    // обратите внимание, что мы оборачиваем уже обёрнутую функцию, а значение счётчика вызовов при этом сохраняется
    countAndMetricPrint := metricTimeCall(countedPrint)
    countAndMetricPrint("Hello")
    countAndMetricPrint("World")

	// Результат
	// Функция main.myprint вызвана 1 раз
	// Hello world
	// Функция main.myprint вызвана 2 раз
	// Hi
	// Функция main.myprint вызвана 3 раз
	// Hello
	// Execution time:  3.147µs
	// Функция main.myprint вызвана 4 раз
	// World
	// Execution time:  3.16µs

}

func PrintAllFiles(path string) {
    // получаем список всех элементов в папке (и файлов, и директорий)
    files, err := os.ReadDir(path)
    if err != nil {
        fmt.Println("unable to get list of files", err)
        return
    }
    //  проходим по списку
    for _, f := range files {
        // получаем имя элемента
        // filepath.Join — функция, которая собирает путь к элементу с разделителями
        filename := filepath.Join(path, f.Name())
        // печатаем имя элемента
        fmt.Println(filename)
        // если элемент — директория, то вызываем для него рекурсивно ту же функцию
        if f.IsDir() {
            PrintAllFiles(filename)
        }
    }
}

// На основе функции PrintAllFiles из предыдущего примера реализуйте функцию PrintAllFilesWithFilter(path string, filter string), 
// которая будет печатать только путь со строкой filter. 
func PrintAllFilesWithFilter(path string, filter string) {
	// получаем список всех элементов в папке (и файлов, и директорий)
    files, err := os.ReadDir(path)
    if err != nil {
        fmt.Println("unable to get list of files", err)
        return
    }
    //  проходим по списку
    for _, f := range files {
        // получаем имя элемента
        // filepath.Join — функция, которая собирает путь к элементу с разделителями
        filename := filepath.Join(path, f.Name())

		if strings.Contains(filename, filter) {
			// печатаем имя элемента
        	fmt.Println(filename)
		}

		// если элемент — директория, то вызываем для него рекурсивно ту же функцию
		if f.IsDir() {
			PrintAllFiles(filename)
		}
    }
}


// В примере этого урока мы рассматривали функцию PrintAllFilesWithFilterClosure. В качестве параметра она принимает обязательную строку из пути файла, имя которого выводится на печать.
// На её основе напишите функцию PrintFilesWithFuncFilter(path string, predicate func (string) bool).
// В качестве второго параметра принимается функция, которая проверяет свой аргумент на соответствие определённому условию. Если оно выполняется, то функция возвращает true.
// Для примера может быть передана такая функция:
// containsDot возвращает все пути, содержащие точки
func containsDot(s string) bool {
    return strings.Contains(s, ".")
}

func PrintFilesWithFuncFilter(path string, predicate func (string) bool) {
	// получаем список всех элементов в папке (и файлов, и директорий)
    files, err := os.ReadDir(path)
    if err != nil {
        fmt.Println("unable to get list of files", err)
        return
    }
    //  проходим по списку
    for _, f := range files {
        // получаем имя элемента
        // filepath.Join — функция, которая собирает путь к элементу с разделителями
        filename := filepath.Join(path, f.Name())

		if (predicate(filename)) {
			// печатаем имя элемента
        	fmt.Println(filename)
		}

		// если элемент — директория, то вызываем для него рекурсивно ту же функцию
		if f.IsDir() {
			PrintAllFiles(filename)
		}
    }
}


// countCall — функция-обёртка для подсчёта вызовов
func countCall(f func(string)) func(string) {
    cnt := 0
    // получаем имя функции. Подробнее об этом вызове будет рассказано в следующем курсе
    funcname := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
    return func(s string) {
        cnt++
        fmt.Printf("Функция %s вызвана %d раз\n", funcname, cnt)
        f(s)
    }
}

// metricTimeCall — функция-обёртка для замера времени
func metricTimeCall(f func(string)) func(string) {
    return func(s string) {
        start := time.Now() // получаем текущее время
        f(s)
        fmt.Println("Execution time: ", time.Since(start)) // получаем интервал времени как разницу между двумя временными метками
    }
}

func myprint(s string) {
    fmt.Println(s)
}

func main() {

}


