package main

import (
	"fmt"
	"strconv"
)

func main() {
	for {
		fmt.Println("Введите операцию (+, -, *, /) или 'exit' для выхода:")
		operator := readOperator()
		if operator == "exit" {
			break
		}

		fmt.Println("Введите первое число:")
		num1 := readNumber()
		fmt.Println("Введите второе число:")
		num2 := readNumber()

		result, err := calculate(num1, operator, num2)
		if err != nil {
			fmt.Println("Ошибка при выполнении операции:", err)
		} else {
			fmt.Printf("Результат: %v\n", result)
		}
	}
}

func readOperator() string {
	var operator string
	fmt.Scanln(&operator)
	return operator
}

func readNumber() float64 {
	var input string
	fmt.Scanln(&input)
	num, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Println("Неверный формат числа.")
		return readNumber()
	}
	return num
}

func calculate(x float64, operator string, y float64) (float64, error) {
	switch operator {
	case "+":
		return x + y, nil
	case "-":
		return x - y, nil
	case "*":
		return x * y, nil
	case "/":
		if y == 0 {
			return 0, fmt.Errorf("Деление на ноль")
		}
		return x / y, nil
	default:
		return 0, fmt.Errorf("Неизвестный оператор: %s", operator)
	}
}
