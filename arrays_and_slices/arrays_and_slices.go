package arrays_and_slices

import "fmt"


func Start() {
	fmt.Println("Работа с массивами и со слайсами");

	weekTempArr := [7]int{1, 2, 3, 4, 5, 6, 7}
    workDaysSlice := weekTempArr[:5] // начнет с начала по 5
    weekendSlice := weekTempArr[5:] // начнет с 6 и до конца
    fromTuesdayToThursDaySlice := weekTempArr[1:4] // начнет с 2 по 4
    weekTempSlice := weekTempArr[:] // возьмет все
    
    fmt.Println(workDaysSlice, len(workDaysSlice), cap(workDaysSlice)) // [1 2 3 4 5] 5 7
    fmt.Println(weekendSlice, len(weekendSlice), cap(weekendSlice)) // [6 7] 2 2 
    fmt.Println(fromTuesdayToThursDaySlice, len(fromTuesdayToThursDaySlice), cap(fromTuesdayToThursDaySlice)) // [2 3 4] 3 6 
    fmt.Println(weekTempSlice, len(weekTempSlice), cap(weekTempSlice)) // [1 2 3 4 5 6 7] 7 7 

	// Попробуйте изменить какой-либо элемент в одном из этих слайсов через операцию []. Убедитесь, что значения изменились и в других слайсах. 
	// Все слайсы в этом примере действительно смотрят на один и тот же базовый массив. Обратите внимание на ёмкость — она равна размеру базового массива, кроме первых элементов.
	// В слайс workDaysSlice взяли элементы с 0 по 4, однако 5 и 6 по-прежнему остаются в массиве. 
	workDaysSlice[3] = 8;

	fmt.Println(workDaysSlice, len(workDaysSlice), cap(workDaysSlice)) // [1 2 3 4 5] 5 7
    fmt.Println(weekendSlice, len(weekendSlice), cap(weekendSlice)) // [6 7] 2 2 
    fmt.Println(fromTuesdayToThursDaySlice, len(fromTuesdayToThursDaySlice), cap(fromTuesdayToThursDaySlice)) // [2 3 4] 3 6 
    fmt.Println(weekTempSlice, len(weekTempSlice), cap(weekTempSlice)) // [1 2 3 4 5 6 7] 7 7 
}