package arrays_and_slices

import (
	"fmt"
	"slices"
)


func Start() {
	fmt.Println("Работа с массивами и со слайсами");

    task8()

	weekTempArr := [7]int{1, 2, 3, 4, 5, 6, 7}
    workDaysSlice := weekTempArr[:5] // начнет с начала по 5
    //weekendSlice := weekTempArr[5:] // начнет с 6 и до конца
    //fromTuesdayToThursDaySlice := weekTempArr[1:4] // начнет с 2 по 4
    //weekTempSlice := weekTempArr[:] // возьмет все
    
    //fmt.Println(workDaysSlice, len(workDaysSlice), cap(workDaysSlice)) // [1 2 3 4 5] 5 7
    //fmt.Println(weekendSlice, len(weekendSlice), cap(weekendSlice)) // [6 7] 2 2 
    //fmt.Println("fromTuesdayToThursDaySlice", fromTuesdayToThursDaySlice, len(fromTuesdayToThursDaySlice), cap(fromTuesdayToThursDaySlice)) // [2 3 4] 3 6 
    //fmt.Println(weekTempSlice, len(weekTempSlice), cap(weekTempSlice)) // [1 2 3 4 5 6 7] 7 7 

	// Попробуйте изменить какой-либо элемент в одном из этих слайсов через операцию []. Убедитесь, что значения изменились и в других слайсах. 
	// Все слайсы в этом примере действительно смотрят на один и тот же базовый массив. Обратите внимание на ёмкость — она равна размеру базового массива, кроме первых элементов.
	// В слайс workDaysSlice взяли элементы с 0 по 4, однако 5 и 6 по-прежнему остаются в массиве. 
	workDaysSlice[3] = 8;

	//fmt.Println(workDaysSlice, len(workDaysSlice), cap(workDaysSlice)) // [1 2 3 4 5] 5 7
    //fmt.Println(weekendSlice, len(weekendSlice), cap(weekendSlice)) // [6 7] 2 2 
    //fmt.Println(fromTuesdayToThursDaySlice, len(fromTuesdayToThursDaySlice), cap(fromTuesdayToThursDaySlice)) // [2 3 4] 3 6 
    //fmt.Println(weekTempSlice, len(weekTempSlice), cap(weekTempSlice)) // [1 2 3 4 5 6 7] 7 7 


    // работа со слайсами
    //a := []int{1, 2, 3, 4} 
    //fmt.Println("a = ", a, len(a), cap(a))
    //b := a[2:3]   // b = [3]
    //fmt.Println("b = ", b, len(b), cap(b))
    //c := append(b, 4, 5, 6,7,8,9,10, 12)
    //fmt.Println("c = ", c, len(c), cap(c))

    //s := make([]int, 4, 7) // [0 0 0 0], len = 4 cap = 7
    // 1. Создаём слайс s с базовым массивом на 7 элементов. 
    // Четыре первых элемента будут доступны в слайсе.

    //slice1 := append(s[:2], 2, 3, 4)  
    //fmt.Println(s, slice1) // [0 0 2 3] [0 0 2 3 4]
    // 2. Берём слайс из первых двух элементов s и добавляем к ним три элемента.
    // Так как суммарная длина полученного слайса (len == 5) меньше ёмкости s[:2] (cap == 7), 
    // то базовый массив остаётся прежним.
    // Слайс s тоже изменился, но его длина осталась прежней

    //slice2 := append(s[1:2], 7) 
    //fmt.Println(s, slice1, slice2) // [0 0 7 3] [0 0 7 3 4] [0 7]
    // 3. Здесь также базовый массив остаётся прежним, изменились все три слайса

    //slice3 := append(s, slice1[1:]...)
    //fmt.Println(len(slice3), cap(slice3))  // 8 14
    // 4. Длина s и slice1[1:] равна 4, длина нового слайса будет равна 8,  
    // что больше ёмкости базового массива.
    // Будет создан новый базовый массив ёмкостью 14,
    // ёмкость нового базового массива подбирается автоматически 
    // и зависит от текущего размера и количества добавленных элементов

    // 5. Легко проверить, что slice3 ссылается на новый базовый массив
    //s[1] = 99
    //fmt.Println(s, slice1, slice2, slice3) 
    // [0 99 7 3] [0 99 7 3 4] [99 7] [0 0 7 3 0 7 3 4] 
}


func task8() {

    fmt.Println("Выполнение задачи 8")
    //  Создайте слайс и заполните его числами от 1 до 100. 
    // Оставьте в слайсе первые и последние 10 элементов и разверните слайс в обратном порядке. 
    // Подумайте, можно ли подобным образом развернуть строку?

    s := make([]int, 0, 100)
    //fmt.Println(s)
    
    for i := 1; i <= 100; i++ {
        s = append(s, i)
    }
    //fmt.Println(s)
    s = append(s[:10], s[90:]...)

    fmt.Println(s)
    slices.Reverse(s)
    fmt.Println(s)


    // // Разернуть строку таким образом не получится, так как строка — неизменяемый тип данных. 
    // Строку можно преобразовать к слайсу байт ([]byte), но это опасно, если строка содержит Unicode-символы
    // Лучше создать слайс рун из строки и развернуть его
    // Например, так:
    input := "The quick brown 狐 jumped over the lazy 犬"
    // Get Unicode code points. 
    n := 0
    // создаём слайс рун 
    runes := make([]rune, len(input))
    // добавляем руны в слайс
    for _, r := range input {
        runes[n] = r
        n++
    }
    // убираем лишние нулевые руны
    runes = runes[0:n]
    // разворачиваем
    for i := 0; i < n/2; i++ {
        runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
    }
    // преобразуем в строку UTF-8. 
    output := string(runes)
    fmt.Println(output)

}