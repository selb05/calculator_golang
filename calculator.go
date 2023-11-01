package main

import (
    "fmt"
    "strconv"//Преобразование строк в числы и наоборот
    "strings"
)

func main() {
    for {
        // Ввод пользователя
        fmt.Print("Введите выражение (или 'exit' для выхода): ")
        input := readInput()

        // Проверка на выход
        if input == "exit" {
            fmt.Println("Выход из программы.")
            break
        }

        // Разбивка введенной строки на числа и оператор
        parts := strings.Fields(input)//Разбивает строку на части,раздленными пробелами и сохраняет их в массив
        if len(parts) != 3 {
            fmt.Println("Неверный формат. Введите выражение в формате 'число оператор число'. Например, '2 + 3'.")
            continue
        }

        // Преобразование чисел из строк в числа
        num1, err1 := strconv.Atoi(parts[0])//Преобразует первую часть введенной строки в целое число
        operator := parts[1]
        num2, err2 := strconv.Atoi(parts[2])

        if err1 != nil || err2 != nil {//Если не равно нулю то значит есть ошибка и принтуем инфу об этом
            fmt.Println("Неверный формат чисел.")
            continue
        }

        // Проверка на диапазон чисел
		if num1 < 1 || num1 > 10 || num2 < 1 || num2 > 10 {
			fmt.Println("Числа должны быть от 1 до 10.")
			continue
		}

        // Выполнение операции
        result, err := calculate(num1, operator, num2)
        if err != nil {
            fmt.Println("Ошибка при выполнении операции:", err)
        } else {
            fmt.Println("Результат:", result)
        }
    }
}

func readInput() string {//Считывает строку из консоли и возвращает её
    var input string
    fmt.Scanln(&input)
    return input
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
