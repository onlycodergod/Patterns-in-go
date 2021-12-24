package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

/*
=== Утилита grep ===
Реализовать утилиту фильтрации (man grep)
Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.

Запуск:
go run task.go sort example.txt
go run task.go -v -n sort example.txt
go run task.go -A=2 sort example.txt
*/

// Основная структура
type Flag struct {
	after      int
	before     int
	context    int
	count      bool
	ignoreCase bool
	invert     bool
	fixed      bool
	lineNum    bool
}

func main() {
	after := flag.Int("A", 0, "печатать +N строк после совпадения")
	before := flag.Int("B", 0, "печатать +N строк до совпадения")
	context := flag.Int("C", 0, "(A+B) печатать ±N строк вокруг совпадения")
	count := flag.Bool("c", false, "количество строк")
	ignoreCase := flag.Bool("i", false, "игнорировать регистр")
	invert := flag.Bool("v", false, "вместо совпадения, исключать")
	fixed := flag.Bool("F", false, "точное совпадение со строкой, не паттерн")
	lineNum := flag.Bool("n", false, "печатать номер строки")

	flag.Parse()           // анализ флагов из командной строки
	pattern := flag.Arg(0) // возврат первого аргумента из командной строки

	var input string
	if fileName := flag.Arg(1); fileName != "" {
		file, err := os.Open(fileName)
		if err != nil {
			fmt.Println("error opening file: err:", err)
			os.Exit(1)
		}
		defer file.Close()

		bytes, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		input = string(bytes)
		input = "sort example.txt"

	} else {
		fmt.Println("Enter text: ")
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString(' ')
		input = text
	}

	parameters := Flag{
		after:      *after,
		before:     *before,
		context:    *context,
		count:      *count,
		ignoreCase: *ignoreCase,
		invert:     *invert,
		fixed:      *fixed,
		lineNum:    *lineNum,
	}

	output, err := grep(input, pattern, parameters)
	if err != nil {
		fmt.Println(err)
	}
	for _, line := range output {
		fmt.Println(line)
	}

}

func grep(input string, pattern string, parameters Flag) ([]string, error) {
	lines := strings.Split(input, "\n")
	var result []string

	if parameters.ignoreCase {
		pattern = "(?i)" + pattern
	}

	if parameters.fixed {
		pattern = "^" + pattern + "$"
	}

	reg, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}

	if parameters.count {
		var count int
		for _, line := range lines {
			if reg.MatchString(line) {
				count++
			}
		}
		num := strconv.Itoa(count)
		result = append(result, num)

		return result, nil
	}

	if parameters.invert {
		for _, line := range lines {
			if !reg.MatchString(line) {
				result = append(result, line)
			}
		}
		return result, nil
	}

	linesAfter := parameters.after
	linesBefore := parameters.before

	if parameters.context != 0 {
		linesAfter = parameters.context
		linesBefore = parameters.context
	}

	var strPosition []int
	var length int
	for i, line := range lines {
		if reg.MatchString(line) {
			if linesAfter == 0 && linesBefore == 0 {
				strPosition = append(strPosition, i)
			} else if linesAfter != 0 {

				if linesAfter+i > len(lines[i:]) {
					length = len(lines[i:]) + i

				} else {
					length = i + linesAfter + 1
				}
				for j := i; j < length; j++ {
					strPosition = append(strPosition, j)
				}
			}
		}
	}

	if linesBefore != 0 {
		for i := len(lines) - 1; i >= 0; i-- {
			if reg.MatchString(lines[i]) {

				if i-linesBefore > 0 {
					length = i - linesBefore
				} else {
					length = 0
				}

				for j := i; j >= length; j-- {
					strPosition = append(strPosition, j)
				}

			}
		}
	}

	for _, position := range removeDuplicateInt(strPosition) {
		var output string

		if parameters.lineNum {
			strPosition := strconv.Itoa(position + 1)
			output = strPosition + ":" + lines[position]
		} else {
			output = lines[position]
		}

		result = append(result, output)
	}

	return result, nil
}

func removeDuplicateInt(intSlice []int) []int {
	allKeys := make(map[int]bool)
	var list []int
	for _, item := range intSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	sort.Ints(list)
	return list
}
