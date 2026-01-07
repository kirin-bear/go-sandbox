package hello_generation

var age int

func GetHello() string {
	switch {
	case age >= 1946 && age <= 1964:
		return "Привет, бумер"
	case age >= 1965 && age <= 1980:
		return "Привет, представитель X!"
	case age >= 1981 && age <= 1996:
		return "Привет, миллениал!"
	case age >= 1997 && age <= 2012:
		return "Привет, зумер!"
	case age >= 2013:
		return "Привет, альфа!"
	default:
		return "Привет, хз, кто ты такой"
	}
}

func SetAge(a int) {
	age = a;
}