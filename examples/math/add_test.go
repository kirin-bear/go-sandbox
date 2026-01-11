package math

import "testing"

func TestAddPositive(t *testing.T) {
    sum, err := Add(1, 2)
    if err != nil {
        t.Error("unexpected error")
    }
    if sum != 3 {
        t.Errorf("sum expected to be 3; got %d", sum)
    }
}


func TestAddNegative(t *testing.T) {
    _, err := Add(-1, 2)
    if err == nil {
        t.Error("first arg negative - expected error not be nil" )
    }
    _, err = Add(1, -2)
    if err == nil {
        t.Error("second arg negative - expected error not be nil" )
    }
    _, err = Add(-1, -2)
    if err == nil {
        t.Error("all arg negative - expected error not be nil" )
    }
}


func TestAddIsZero(t *testing.T) {
    _, err := Add(0, 2)
    if err == nil {
        t.Error("first arg is zero" )
    }
    _, err = Add(1, 0)
    if err == nil {
        t.Error("second arg is zero" )
    }

	_, err = Add(0, 0)
    if err == nil {
        t.Error("first and second args are zero" )
    }
}

func TestAdd(t *testing.T) {
    // args — описывает аргументы тестируемой функции
    type args struct {
        a int
        b int
    }
    // описывает структуру тестовых данных и сами тесты
    tests := []struct {
        name    string // название теста
        args    args // аргументы 
        want    int // ожидаемое значение
        wantErr bool // должна ли функция вернуть ошибку 
    }{
         {
            name: "Test Positive",
            args: args{
                a: 1,
                b: 2,
            },
            want:    3,
            wantErr: false,
        },
        {
            name: "Test Negative 1",
            args: args{
                a: -1,
                b: 2,
            },
            want:    0,
            wantErr: true,
        },
        {
            name: "Test Negative 2",
            args: args{
                a: 1,
                b: -2,
            },
            want:    0,
            wantErr: true,
        },

        {
            name: "Test Negative all",
            args: args{
                a: -1,
                b: -2,
            },
            want:    0,
            wantErr: true,
        }, 
    }
    // вызываем тестируемую функцию для каждого тестового случая  
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := Add(tt.args.a, tt.args.b)
            if (err != nil) != tt.wantErr {
                t.Errorf("Add() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if got != tt.want {
                t.Errorf("Add() = %v, want %v", got, tt.want)
            }
        })
    }
}
 