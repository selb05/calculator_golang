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

		fmt.Println("Введите первое целое число:")
		num1 := readInteger()
		fmt.Println("Введите второе целое число:")
		num2 := readInteger()

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
    if operator != "+" && operator != "-" && operator != "*" && operator != "/" {
        fmt.Println("Неверный оператор. Допустимые операторы: +, -, *, /")
        return readOperator()
    }
    return operator
}

func readInteger() int {
	var input string
	fmt.Scanln(&input)
	num, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Неверный формат числа. Введите целое число")
		return readInteger()
	}
	return num
}

func calculate(x int, operator string, y int) (int, error) {
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
