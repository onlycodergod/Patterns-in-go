package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
=== Утилита cut ===
Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные
Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.

Запуск:
echo "some:text:for:example" | go run task.go -d=':' -f=1-2
echo "some:text:for:example" | go run task.go  -f=1 -s
*/

func main() {
	conf := NewConfig()
	Start(conf)
}

// Config - конфигурация программы
type Config struct {
	separated bool
	delim     string
	fields    string
}

// NewConfig - конструктор, парсящий флаги и аргументы
func NewConfig() *Config {
	conf := Config{}

	flagS := flag.Bool("s", false, "только строки с разделителем")
	flag.StringVar(&conf.delim, "d", "", "использовать другой разделитель")
	flag.StringVar(&conf.fields, "f", "", "выбрать поля (колонки)")
	flag.Parse()

	conf.separated = *flagS
	return &conf
}

func Start(conf *Config) {
	var str strings.Builder
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		str.WriteString(sc.Text())
	}

	result, err := cut(str.String(), conf)
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Print(result)
}

func cut(row string, conf *Config) (string, error) {
	var result strings.Builder
	fields := make(map[int]bool)

	// разделитель
	delim := "\t"
	if conf.delim != "" {
		if len(conf.delim) == 1 {
			delim = conf.delim
		} else {
			return "", fmt.Errorf("err: разделитель должен быть одним символом ")
		}
	}

	// колонки
	if conf.fields != "" {
		sequence := strings.Split(conf.fields, ",")

		for _, seqPart := range sequence {
			seqPartRange := strings.Split(strings.TrimSpace(seqPart), "-")
			if len(seqPartRange) == 2 {
				seqPartRangeNumber1, err := strconv.Atoi(seqPartRange[0])
				if err != nil {
					return "", fmt.Errorf("err: недопустимое значение поля: '%s'", seqPartRange[0])
				}

				seqPartRangeNumber2, err := strconv.Atoi(seqPartRange[1])
				if err != nil {
					return "", fmt.Errorf("err: недопустимое значение поля: '%s'", seqPartRange[1])
				}

				if seqPartRangeNumber1 > seqPartRangeNumber2 {
					return "", fmt.Errorf("err: недопустимый диапазон")
				}

				if seqPartRangeNumber1 < 1 {
					return "", fmt.Errorf("err: поля нумеруются от 1")
				}

				for i := seqPartRangeNumber1; i <= seqPartRangeNumber2; i++ {
					fields[i] = true
				}
			} else {
				fieldNum, err := strconv.Atoi(strings.TrimSpace(seqPart))
				if err != nil {
					return "", fmt.Errorf("err: недопустимое значение поля: '%s'", seqPart)
				}

				if fieldNum < 1 {
					return "", fmt.Errorf("err: поля нумеруются от 1")
				}

				fields[fieldNum] = true
			}
		}
	} else {
		return "", fmt.Errorf("err: необходимо указать список байтов, символов или полей ")
	}

	splittedRow := strings.Split(row, delim)
	if conf.separated && len(splittedRow) == 1 {
		return "", nil
	}

	isNeedDelim := false
	for i, part := range splittedRow {
		_, ok := fields[i+1]
		if ok {
			if isNeedDelim {
				result.WriteString(delim + part)
			} else {
				result.WriteString(part)
				isNeedDelim = true
			}
		}
	}
	//fmt.Printf("%#v", splittedRow)
	return result.String(), nil
}
