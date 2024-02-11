package main

import (
	"fmt"
	"math/big"
)

func main() {
	// Объявление и инициализация переменных a и b
	a := big.NewFloat(22222222222.111)
	b := big.NewFloat(33333333333.1234)

	// Сложение
	sum := new(big.Float)
	fmt.Printf("%f + %f = ", a, b)
	fmt.Println(sum.Add(a, b))
	// Вычитание
	substraction := new(big.Float)
	fmt.Printf("%f - %f = ", a, b)
	fmt.Println(substraction.Sub(a, b))

	// Умножение
	multiplication := new(big.Float)
	fmt.Printf("%f / %f = ", a, b)
	fmt.Println(multiplication.Mul(a, b))

	// Деление
	division := new(big.Float)
	fmt.Printf("%f / %f = ", a, b)
	fmt.Println(division.Quo(a, b))
}
