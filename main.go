package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Массив с римскими числами
var roman = [101]string{"O", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X", "XI", "XII", "XIII", "XIV", "XV", "XVI", "XVII", "XVIII", "XIX", "XX",
	"XXI", "XXII", "XXIII", "XXIV", "XXV", "XXVI", "XXVII", "XXVIII", "XXIX", "XXX", "XXXI", "XXXII", "XXXIII", "XXXIV", "XXXV", "XXXVI", "XXXVII", "XXXVIII", "XXXIX", "XL",
	"XLI", "XLII", "XLIII", "XLIV", "XLV", "XLVI", "XLVII", "XLVIII", "XLIX", "L", "LI", "LII", "LIII", "LIV", "LV", "LVI", "LVII", "LVIII", "LIX", "LX",
	"LXI", "LXII", "LXIII", "LXIV", "LXV", "LXVI", "LXVII", "LXVIII", "LXIX", "LXX", "LXXI", "LXXII", "LXXIII", "LXXIV", "LXXV", "LXXVI", "LXXVII", "LXXVIII", "LXXIX", "LXXX",
	"LXXXI", "LXXXII", "LXXXIII", "LXXXIV", "LXXXV", "LXXXVI", "LXXXVII", "LXXXVIII", "LXXXIX", "XC", "XCI", "XCII", "XCIII", "XCIV", "XCV", "XCVI", "XCVII", "XCVIII", "XCIX", "C"}

// Функция main
func main() {

	var line string

	fmt.Println("Введите выражение (арабскими либо римскими числами):")
	//Чтение строки
	reader := bufio.NewReader(os.Stdin)
	line, _ = reader.ReadString('\n')
	line = strings.TrimSuffix(line, "\n")

	fmt.Println(parseLine(line))
}

// Функция, в которой содержатся все основные вызовы
func parseLine(line string) string {

	var number1 int
	var number2 int
	var romanNumb bool
	var result string

	//Прасинг строки через сепаратор
	separator := func(c rune) bool {
		return c == ' ' || c == '+' || c == '-' || c == '*' || c == '/'
	}
	numbers := strings.FieldsFunc(line, separator)

	//Проврека на правильность ввода(два операнда)
	if len(numbers) < 2 {
		panic("Вывод ошибки, так как строка не является математической операцией.")
	} else if len(numbers) > 2 {
		panic("Вывод ошибки, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
	}

	//Проверка оператора
	operation := findOperation(line)
	if operation == "null" {
		panic("Вывод ошибки, строка введена некорректно")
	}

	//Ковертация и проверка на правильность введенного формата
	if isRoman(numbers[0]) && isRoman(numbers[1]) {
		number1 = convertToArabian(numbers[0])
		number2 = convertToArabian(numbers[1])
		romanNumb = true
	} else if !isRoman(numbers[0]) && !isRoman(numbers[1]) {
		number1, _ = strconv.Atoi(numbers[0])
		number2, _ = strconv.Atoi(numbers[1])
	} else {
		panic("Вывод ошибки, так как используются одновременно разные системы счисления.\n")
	}

	//Проверка на введенные числа
	if number1 < 1 || number1 > 10 || number2 < 1 || number2 > 10 {
		panic("Вывод ошибки, так как числа должны быть от 1 до 10")
	}

	//Вычисление
	resultInt := calculate(number1, number2, operation)

	//Проверка на римское число и вычисление результата
	if romanNumb {
		if resultInt <= 0 {
			panic("Вывод ошибки, так как римское число должно быть больше нуля")
		} else {
			result = convertToRoman(resultInt)
		}
	} else {
		result = strconv.Itoa(resultInt)
	}
	return result
}

// Функция, которая проверяет о оператор
func findOperation(line string) string {
	var oper string = "null"
	if strings.Contains(line, "+") {
		oper = "+"
	}
	if strings.Contains(line, "-") {
		oper = "-"
	}
	if strings.Contains(line, "/") {
		oper = "/"
	}
	if strings.Contains(line, "*") {
		oper = "*"
	}
	return oper
}

// Функция, которая проверяет в римском ли формате введено число
func isRoman(numb string) bool {

	ret := false

	for i := 0; i < 101; i++ {
		if numb == roman[i] {
			ret = true
			break
		}
	}
	return ret
}

// Функция, которая конвертирует число в арабское
func convertToArabian(number string) int {

	var convertedNumb int
	for i := 0; i < len(roman); i++ {
		if number == roman[i] {
			convertedNumb = i
			break
		}
	}
	return convertedNumb
}

// Функция, которая вычисляет ответ
func calculate(numb1 int, numb2 int, operator string) int {

	var result int

	if operator == "+" {
		result = numb1 + numb2
	} else if operator == "-" {
		result = numb1 - numb2
	} else if operator == "/" {
		result = numb1 / numb2
	} else if operator == "*" {
		result = numb1 * numb2
	}

	return result
}

// Функция, которая конвертирует в римское число
func convertToRoman(numb int) string {
	return roman[numb]
}
