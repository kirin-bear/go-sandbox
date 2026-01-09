package structures

import (
	"encoding/json"
	"fmt"
	"strings"
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
	Name        string    `json:"Имя"`
	Email       string    `json:"Почта"`
	DateOfBirth time.Time `json:"-"` // убираем из преобразования в json
}

// NewPrivateFoo — конструктор типа privateFoo
// Функция публичная, то есть может быть вызвана из других пакетов
func NewPrivateFoo() privateFoo {
	return privateFoo{Value: "some data"}
}

func Start() {
	//task1()
	task2()
}

// Задача 1.
// Для сериализации используется функция json.Marshal() пакета json. Дана структура: Person
// Напишите код, который будет сериализовывать структуру в json-строку следующего вида:
func task1() {

	p := Person{
		Name:  "Алекс",
		Email: "alex@yandex.ru",
	}
	j, err := json.Marshal(p)
	if err != nil {
		fmt.Println("Unable marshal to json")
	}
	fmt.Println(string(j))
}

// Задача 2.
//
// Есть пример API-вызова в формате JSON:
//
//	{
//	    "header": {
//	        "code": 0,
//	        "message": ""
//	    },
//	    "data": [{
//	        "type": "user",
//	        "id": 100,
//	        "attributes": {
//	            "email": "bob@yandex.ru",
//	            "article_ids": [10, 11, 12]
//	        }
//	    }]
//	}
//
// К сожалению, ни Swagger-описания, ни статьи с API-ответом в любимом сервисе заметок — нет.
// Опишите данный объект в виде структуры на Go, в учебных целях отбросив «так делать нельзя» и «это не дело».
// На входе есть строка с сырыми данными, требуется написать функцию её десериализации:
type Response struct {
	// поля с тегами
	Header Header `json:"header"`
	Data   []DataItem `json:"data"`
}

type Header struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type DataItem struct {
	Type string `json:"type"`
	Id int `json:"id"`
	Attributes Attributes `json:"attributes"`
}

type Attributes struct {
	Email      string `json:"email"`
	ArticleIds []int `json:"article_ids"`
}

func readResponse(rawResp string) Response {
	// код десереализации
	return Response{}
}

func task2() {

	prettyJson := `{
		"header": {
			"code": 0,
			"message": "Hello"
		},
		"data": [{
			"type": "user",
			"id": 100,
			"attributes": {
				"email": "bob@yandex.ru",
				"article_ids": [10, 11, 12]
			}
		}]
	}`

	// Удаляем все пробельные символы
	compactJson := strings.ReplaceAll(prettyJson, " ", "")  // Удалит пробелы
	compactJson = strings.ReplaceAll(compactJson, "\n", "") // Удалит переносы
	compactJson = strings.ReplaceAll(compactJson, "\t", "") // Удалит табуляции
	fmt.Println(compactJson)

	r := Response{}

	err := json.Unmarshal([]byte(compactJson), &r)

	if err != nil {
		fmt.Println("Error Unmarshall", err)
	}

	fmt.Printf("%+v\n", r)

}
