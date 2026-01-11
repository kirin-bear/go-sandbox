package estimatedtst

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestEstimateV(t *testing.T) {

    // описывает структуру тестовых данных и сами тесты
    tests := []struct {
        name    string // название теста
        input   int // входной аргумент
        want    string // ожидаемое значение
        wantErr bool // должна ли функция вернуть ошибку 
    }{
        {
            name: "Test Less Than 10 is small (1)",
            input: 1,
            want: "small",
        },
        {
            name: "Test Less Than 10 is small (-1)",
            input: -1,
            want: "small",
        },
        {
            name: "Test Less Than 10 is small (9)",
            input: 9,
            want: "small",
        },
        {
            name: "Test Less Than 100 is medium (10)",
            input: 10,
            want: "medium",
        },
        {
            name: "Test Less Than 100 is medium (99)",
            input: 99,
            want: "medium",
        },
        {
            name: "Test is big (100)",
            input: 100,
            want: "big",
        },
    }
    // вызываем тестируемую функцию для каждого тестового случая  
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := EstimateValue(tt.input)
            if got != tt.want {
                t.Errorf("Add() = %v, want %v", got, tt.want)
            }
        })
    }


    // или

    t.Run("Small", func(t *testing.T) {
        assert.Equal(t, "small", EstimateValue(9))
    })

    t.Run("Medium", func(t *testing.T) {
        assert.Equal(t, "medium", EstimateValue(99))
    })

    t.Run("Big", func(t *testing.T) {
        assert.Equal(t, "big", EstimateValue(100))
    })

}


func TestEstimateValueTableDriven(t *testing.T) {
    testCases := []struct {
        Name          string
        InputValue    int
        ExpectedValue string
    }{
        {
            Name:          "Small",
            InputValue:    9,
            ExpectedValue: "small",
        },
        {
            Name:          "Medium",
            InputValue:    99,
            ExpectedValue: "medium",
        },
        {
            Name:          "Big",
            InputValue:    100,
            ExpectedValue: "big",
        },
    }

    for _, tc := range testCases {
        t.Run(tc.Name, func(t *testing.T) {
            assert.EqualValues(t, tc.ExpectedValue, EstimateValue(tc.InputValue))
        })
    }
} 