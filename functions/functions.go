package functions

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)


func Start() {
	fmt.Println("Работа с функциями")

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

