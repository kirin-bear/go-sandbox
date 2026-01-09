package structures

import (
	"encoding/json"
	"fmt"
	"time"
)

// privateFoo — неэкспортируемый тип
type privateFoo struct {
	Value string
}

type GetUserRequest struct {
	UserId    string `json:"user_id" yaml:"user_id" format:"uuid" example:"2e263a90-b74b-11eb-8529-0242ac130003"`
	IsDeleted *bool  `json:"is_deleted,omitempty" yaml:"is_deleted"`
}

type Person struct {
	Name        string `json:"Имя"`
	Email       string `json:"Почта"`
	DateOfBirth time.Time `json:"-"` // убираем из преобразования в json
}

// NewPrivateFoo — конструктор типа privateFoo
// Функция публичная, то есть может быть вызвана из других пакетов
func NewPrivateFoo() privateFoo {
	return privateFoo{Value: "some data"}
}

func Start() {
	// Для сериализации используется функция json.Marshal() пакета json. Дана структура: Person
	// Напишите код, который будет сериализовывать структуру в json-строку следующего вида:

	p := Person{
		Name:  "Алекс",
		Email: "alex@yandex.ru",
	}
	j, err := json.Marshal(p)
	if (err != nil) {
		fmt.Println("Unable marshal to json")
	}
	fmt.Println(string(j))
}
