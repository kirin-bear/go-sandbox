package figures

import (
	"fmt"
	"math"
)

type figures int

const (
    square figures = iota // квадрат
    circle // круг
    triangle // равносторонний треугольник
) 

func area(figure figures) (func(float64) float64, bool) {
	switch figure {
	case square: 
		return areaSquare, true
	case triangle:
		return areaTriangle, true
	case circle:
		return areaCircle, true
	default:
		return nil, false
	}
}

func areaSquare(s float64) float64 {
	return s * s
}

func areaCircle(s float64) float64 {
	return 3.14 * s * s
}

func areaTriangle(s float64) float64 {
	return (s * s * math.Pow(3.0, 1.0/2.0))/4
}

func Start() {
	myFigure := circle
	x:= 2.0
	ar, ok := area(myFigure)
	if !ok {
		fmt.Println("Ошибка")
		return
	}
	myArea := ar(x) 
	fmt.Println(myArea)
}