package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

/*
=== Задача на распаковку ===
Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)
В случае, если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.
Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

/* Запуск тестов:
go test -v

go test -cover
go test -coverprofile=profile.out
go tool cover -html=profile.out
*/

func main() {
	data := []string{"a4bc2d5e", "abcd", "45", "", `qwe\4\5`, `qwe\45`, `qwe\\\\5`}
	for _, str := range data {
		res, err := Unpacking(str)
		fmt.Println(str + " => " + res)
		if err != nil {
			fmt.Println(err)
		}
	}
}

// Unpacking - распаковка строки
func Unpacking(str string) (string, error) {
	// если строка пуста или начинается с числа - можно не продолжать
	if str == "" || unicode.IsDigit([]rune(str)[0]) {
		errStr := "err: некорректная строка"
		return "", fmt.Errorf(errStr)
	}

	const escapeSymbol = 92 // символ escape - последовательности "/"
	var resultStr strings.Builder
	runes := []rune(str)

	for i := 0; i < len(runes); i++ {
		if runes[i] == escapeSymbol {
			i++
			resultStr.WriteRune(runes[i])
		} else if repeats, err := strconv.Atoi(string(runes[i])); err == nil { // находим цифры - повторы
			// предыдущий не нулевой символ тоже цифра (две цифры подряд)
			if i > 0 && unicode.IsDigit(runes[i-1]) && (i > 1 && runes[i-2] != escapeSymbol) {
				return "", fmt.Errorf("invalid rune")
			}
			resultStr.WriteString(strings.Repeat(string(runes[i-1]), repeats-1))
		} else {
			resultStr.WriteRune(runes[i])
		}
	}
	return resultStr.String(), nil
}
