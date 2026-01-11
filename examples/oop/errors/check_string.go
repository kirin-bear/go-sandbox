// Напишите функцию MyCheck(input string) error, которая должна проверять строку на соответствие следующим параметрам:
// Строка не должна содержать цифр (found numbers).
// Длина должна быть меньше 20 символов (line is too long).
// Строка должна иметь два пробела (no two spaces).
// В скобках приведены тексты возвращаемых ошибок.
// Функция должна находить все возможные ошибки за один вызов. Результат должен содержать слайс ошибок, по которым строка не прошла проверку, или быть nil. Подсказка: определите свой тип для слайса ошибок.

package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"unicode"
)

type SliceStringErrors struct {
	prefix string
	errors []error
}

func (sse *SliceStringErrors) AddError(se error) {
	sse.errors = append(sse.errors, se)
}

func (sse *SliceStringErrors) HasErrors() bool {
	return len(sse.errors) > 0
}

// NewTimeError возвращает переменную типа TimeError c текущим временем.
func NewSliceStringError(prefix string) SliceStringErrors {
    return SliceStringErrors{
        prefix: prefix,
    }
}

func (sse SliceStringErrors) Error() string {
	textErrors := strings.Builder{}
    for _, v := range sse.errors {
		textErrors.WriteString(v.Error() + " // ")
	}

	return textErrors.String()
}


// 1) вставьте определение типа для []error
// 2) определите метод Error для вашего типа, который будет выводить
//    все ошибки слайса
// 3) реализуйте функцию MyCheck
//
// ...
func MyCheck(input string) error {

	stringErrors := NewSliceStringError("check_strings")

	// 1. Проверка длины (меньше 20 символов)
	if len(input) >= 20 {
		stringErrors.AddError(errors.New("line is too long"))
	}

	// 2. Проверка на наличие цифр
	for _, r := range input {
		if unicode.IsDigit(r) {
			stringErrors.AddError(errors.New("found numbers"))
			continue
		}
	}

	// 3. Проверка количества пробелов (ровно два)
	if strings.Count(input, " ") != 2 {
		stringErrors.AddError(errors.New("no two spaces"))
	}

	if stringErrors.HasErrors() {
		return stringErrors
	}

	return nil // Все проверки пройдены
}

func main() {
    for {
        fmt.Printf("Укажите строку (q для выхода): ")
        reader := bufio.NewReader(os.Stdin)
        ret, err := reader.ReadString('\n')
        if err != nil {
            fmt.Println(err)
            continue
        }
        ret = strings.TrimRight(ret, "\n")
        if ret == `q` {
            break
        }
        if err = MyCheck(ret); err != nil {
            fmt.Println(err)
        } else {
            fmt.Println(`Строка прошла проверку`)
        }
    }
}


// или так
// type SliceError []error
// 
// func (errs SliceError) Error() string {
//     var out string
//     for _, err := range errs {
//         out += err.Error() + ";"
//     }
//     return strings.TrimRight(out, `;`)
// }
// 
// func MyCheck(input string) error {
//     var (
//         err      SliceError
//         spaces   int
//         hasDigit bool
//     )
//     if len([]rune(input)) >= 20 {
//         err = SliceError{errors.New(`line is too long`)}
//     }
//     for _, ch := range input {
//         if ch == ' ' {
//             spaces++
//         } else if ch >= '0' && ch <= '9' {
//             hasDigit = true
//         }
//     }
//     if hasDigit {
//         err = append(err, errors.New(`found numbers`))
//     }
//     if spaces != 2 {
//         err = append(err, errors.New(`no two spaces`))
//     }
//     if len(err) == 0 {
//         return nil
//     }
//     return err
// } 