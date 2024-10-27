package main

import (
	"errors"
	"fmt"
)

// Задание3
// Функция, которая возвращает приветственное сообщение
func hello(name string) string {
	return "Привет, " + name + "!"
}

// Функция, которая выводит все четные числа в диапазоне
func printEven(a, b int64) error {
	// Проверка диапазона: если левая граница больше правой, возвращаем ошибку
	if a > b {
		return errors.New("левая граница больше правой")
	}

	// Выводим четные числа в диапазоне [a, b]
	for i := a; i <= b; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	// Если ошибок нет, возвращаем nil
	return nil
}

// Функция, которая выполняет арифметическое действие с двумя числами
func apply(a, b float64, operator string) (float64, error) {
	switch operator {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		// Проверяем деление на ноль
		if b == 0 {
			return 0, errors.New("деление на ноль")
		}
		return a / b, nil
	default:
		return 0, errors.New("действие не поддерживается")
	}
}

func main() {
	fmt.Println("Hello, World!")

	// Тестирование функции hello
	fmt.Println(hello("Иван"))

	// Тестирование функции printEven
	fmt.Println("Четные числа в диапазоне [2, 10]:")
	err := printEven(2, 10)
	if err != nil {
		fmt.Println("Ошибка:", err)
	}

	// Попытка вызвать printEven с неправильным диапазоном
	fmt.Println("Четные числа в диапазоне [10, 2]:")
	err = printEven(10, 2)
	if err != nil {
		fmt.Println("Ошибка:", err)
	}

	// Тестирование функции apply
	fmt.Println("Результаты работы функции apply:")
	result, err := apply(3, 5, "+")
	if err == nil {
		fmt.Printf("3 + 5 = %.2f\n", result)
	} else {
		fmt.Println("Ошибка:", err)
	}

	result, err = apply(7, 10, "*")
	if err == nil {
		fmt.Printf("7 * 10 = %.2f\n", result)
	} else {
		fmt.Println("Ошибка:", err)
	}

	result, err = apply(3, 5, "#")
	if err == nil {
		fmt.Printf("Результат = %.2f\n", result)
	} else {
		fmt.Println("Ошибка:", err)
	}

	result, err = apply(10, 0, "/")
	if err == nil {
		fmt.Printf("10 / 0 = %.2f\n", result)
	} else {
		fmt.Println("Ошибка:", err)
	}
}
