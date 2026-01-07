package scope_export

import "fmt"

const Email = "a@b.com" // глобальная экспортируемая константа

var support string // глобальная неэкспортируемая переменная

func SetSupport(s string) {
	support = s
}

// GetContact возвращает имя и email.
func GetContact() string { // экспортируемая функция
    return fmt.Sprintf("%s <%s>", support, Email)
} 