package mapas

import (
	"fmt"
	"strconv"
)


func Start() {

	fmt.Println("Работа с мапами");

	//task1()
	task2()


	input := []string{
		"cat",
		"dog",
		"bird",
		"dog",
		"parrot",
		"cat",
	} 
	fmt.Println(RemoveDuplicates(input))
}


func task1() {
	// Сделайте map для хранения прейскуранта в рублях:
	// хлеб — 50,
	// молоко — 100,
	// масло — 200,
	// колбаса — 500,
	// соль — 20,
	// огурцы — 200,
	// сыр — 600,
	// ветчина — 700,
	// буженина — 900,
	// помидоры — 250,
	// рыба — 300,
	// хамон — 1500.
	// Выведите перечень деликатесов — продуктов дороже 500 рублей.
	// Заказ выдан слайсом []string{"хлеб", "буженина", "сыр", "огурцы"}. Посчитайте стоимость заказа.

	p := map[string]int{
		"хлеб": 50,
		"молоко": 100,
		"масло": 200,
		"колбаса": 500,
		"соль": 20,
		"огурцы": 200,
		"сыр": 600,
		"ветчина": 700,
		"буженина": 900,
		"помидоры": 250,
		"рыба": 300,
		"хамон": 1500,
	}
	deliciouses := make([]string, 0, 12)
	order := []string{"хлеб", "буженина", "сыр", "огурцы"}
	costOrder := 0

	for k, v := range p {
		if (v > 500) {
			deliciouses = append(deliciouses, k)
		}
	}

	for _, position := range order {
		cost, ok := p[position]
		if (ok == true) {
			costOrder+= cost
		}
	}

	fmt.Println(p)
	fmt.Println(deliciouses)
	fmt.Println(costOrder)
}

func task2() {

	arr := []int{1,2,3,4,5,6,7,8,9,10}
	val := 10
	res := []string{}

	for k, v := range arr {
		
		for k2, v2 := range arr {

			if (k2 == k) {
				continue
			}

			if (v + v2 == val) {
				res = append(res, strconv.Itoa(k+1)+"_"+strconv.Itoa(k2+1))
			}

		}

	}

	fmt.Println(res)

	// а как надо
	//func find(arr []int, k int) []int {
    //    // Создадим пустую map  
    //    m := make(map[int]int)
    //    // будем складывать в неё индексы массива, а в качестве ключей использовать само значение 
    //    for i, a := range arr {
    //        if j, ok := m[k-a]; ok { // если значение k-a уже есть в массиве, значит, arr[j] + arr[i] = k и мы нашли, то что нужно
    //            return []int{i, j}
    //        }
    //        // если искомого значения нет, то добавляем текущий индекс и значение в map
    //        m[a] = i
    //    }
    //    // не нашли пары подходящих чисел
    //    return nil
    //}    
    // как можно заметить, алгоритм пройдётся по массиву всего один раз
    // если бы мы искали подходящее значение каждый раз через перебор массива, то пришлось бы сделать гораздо больше вычислений 

}

func RemoveDuplicates(input []string) []string {

	res := []string{}
	m := make(map[string]int)
	

	for k, v := range input {

		_, exists := m[v]
		if (exists) {
			continue
		}

		m[v] = k
		res = append(res, v)

	}

	return res;


	// а как надо было
	// func RemoveDuplicates(input []string) []string {
	//	output := make([]string,0)
	//	inputSet := make(map[string]struct{}, len(input))
	//	for _, v := range input {
	//		if _, ok := inputSet[v]; !ok {
	//			output = append(output, v)
	//
	//		}
	//		inputSet[v] = struct{}{}
	//	}
	//
	//	return output
	//}
}