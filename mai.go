package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	for {
		var reader = bufio.NewReader(os.Stdin)

		fmt.Println("\nВведите пример:")
		var input, _ = reader.ReadString('\n')
		input = strings.TrimSpace(input) // срез пробелов с обоих краёв
		input = strings.ToUpper(input)   // верхний регистр
		var list = strings.Split(input, " ")
		if len(list) != 3 {
			fmt.Println("Числа и математическая операция должны быть разделены пробелом")
			continue
		}
		var firstValue = list[0]
		var operation = list[1]
		var secondValue = list[2]

		if _, err := strconv.Atoi(firstValue); err != nil {
			var number1 = numToArabic(firstValue)
			var number2 = numToArabic(secondValue)
			if counting(number1, number2, operation) < 1 {
				fmt.Print("Результатом не может быть отрицательное число")
				continue
			}
			fmt.Print(numToRoman(counting(number1, number2, operation)))
		} else {
			number1, _ := strconv.Atoi(firstValue)
			if _, err := strconv.Atoi(secondValue); err == nil {
				var number2, _ = strconv.Atoi(secondValue)
				if number1 > 10 || number1 < 0 || number2 > 10 || number2 < 0 {
					fmt.Print("Калькулятором принимаются два числа на вход в диапазоне от 0 до 10. Повторите попытку.")
					continue
				}
				fmt.Print(counting(number1, number2, operation))
			} else {
				fmt.Print("На вход принимаются арабские или римские числа")
				os.Exit(0)
			}
		}
	}
}
func counting(num1, num2 int, sign string) int {
	switch sign {
	case "+":
		return num1 + num2
	case "-":
		return num1 - num2
	case "*":
		return num1 * num2
	case "/":
		if num2 == 0 {
			fmt.Println("На ноль делить нельзя")
			os.Exit(0)
		}
		return num1 / num2
	default:
		fmt.Print("Выбрана неверная операция ")
		os.Exit(0)
	}
	return 0
}
func numToArabic(b string) int {

	switch b {
	case "I":
		return 1
	case "II":
		return 2
	case "III":
		return 3
	case "IV":
		return 4
	case "V":
		return 5
	case "VI":
		return 6
	case "VII":
		return 7
	case "VIII":
		return 8
	case "IX":
		return 9
	case "X":
		return 10
	default:
		fmt.Print("Некорректный ввод")
		os.Exit(0)
	}
	return 0
}

func numToRoman(num int) string {
	var numbers = [101]string{"O", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X",
		"XI", "XII", "XIII", "XIV", "XV", "XVI", "XVII", "XVIII", "XIX", "XX",
		"XXI", "XXII", "XXIII", "XXIV", "XXV", "XXVI", "XXVII", "XXVIII", "XXIX", "XXX",
		"XXXI", "XXXII", "XXXIII", "XXXIV", "XXXV", "XXXVI", "XXXVII", "XXXVIII", "XXXIX", "XL",
		"XLI", "XLII", "XLIII", "XLIV", "XLV", "XLVI", "XLVII", "XLVIII", "XLIX",
		"L", "LI", "LII", "LIII", "LIV", "LV", "LVI", "LVII", "LVIII", "LIX", "LX",
		"LXI", "LXII", "LXIII", "LXIV", "LXV", "LXVI", "LXVII", "LXVIII", "LXIX", "LXX",
		"LXXI", "LXXII", "LXXIII", "LXXIV", "LXXV", "LXXVI", "LXXVII", "LXXVIII", "LXXIX", "LXXX",
		"LXXXI", "LXXXII", "LXXXIII", "LXXXIV", "LXXXV", "LXXXVI", "LXXXVII", "LXXXVIII", "LXXXIX", "XC",
		"XCI", "XCII", "XCIII", "XCIV", "XCV", "XCVI", "XCVII", "XCVIII", "XCIX", "C"}
	return numbers[num]
}
