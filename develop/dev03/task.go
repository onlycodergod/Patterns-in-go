package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

/*
=== Утилита sort ===
Отсортировать строки (man sort)
Основное
Поддержать ключи
-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки
Дополнительное
Поддержать ключи
-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

/*
Запуск тестов:
go test -v

go test -cover
go test -coverprofile=profile.out
go tool cover -html=profile.out

go vet -c=10 -json task.go

Запуск программы:
go run task.go -u 1_sort.txt
*/

func main() {
	s := NewSortConfig()
	res, err := Start(s)
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println(res)
}

// SortConfig - Конфигурация сортировки
type SortConfig struct {
	sortColumn         int
	sortByNumericValue bool
	reverseSort        bool
	uniqueRows         bool
	isAlreadySorted    bool
	filename           string
}

// NewSortConfig - Конструктор конфига
func NewSortConfig() *SortConfig {
	s := SortConfig{}
	flag.IntVar(&s.sortColumn, "k", 0, "Sets column for sort")
	flagN := flag.Bool("n", false, "Makes sort by numeric value")
	flagR := flag.Bool("r", false, "Makes reverse sort")
	flagU := flag.Bool("u", false, "Ignore duplicate lines")
	flagC := flag.Bool("c", false, "Check if rows already sorted")

	flag.Parse()
	args := flag.Args()
	s.sortByNumericValue = *flagN
	s.reverseSort = *flagR
	s.uniqueRows = *flagU
	s.isAlreadySorted = *flagC

	if len(args) == 1 {
		s.filename = args[0]
	} else {
		log.Fatalf("The argument (path to the file name) must be one")
	}

	return &s
}

// Start - Точка входа в программу сортировки
func Start(s *SortConfig) (string, error) {
	rows, err := readFile(s.filename)
	if err != nil {
		return "", fmt.Errorf("can not read file '%s': %s", s.filename, err.Error())
	}
	return sortRows(rows, s)
}

func readFile(filename string) ([]string, error) {
	var rows []string
	file, err := os.Open(filename)
	if err != nil {
		return rows, err
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		rows = append(rows, sc.Text())
	}
	return rows, nil
}

func uniqualizer(rows []string) []string {
	tempBuf := make(map[string]bool)
	for _, row := range rows {
		tempBuf[row] = true
	}

	// "очищаем" исходный слайс и заполняем его снова
	rows = rows[:0]
	for key := range tempBuf {
		rows = append(rows, key)
	}
	return rows
}

func getColumnValue(row string, s *SortConfig) (string, error) {
	// поиск разделителя: один или более пробельных символов
	re := regexp.MustCompile(`\s+`)
	// -b тег по умолчанию - при работе с колонками отрезаем пробелы вначале и в конце.
	listOfColumns := re.Split(strings.TrimSpace(row), -1)
	if len(listOfColumns) >= s.sortColumn {
		return listOfColumns[s.sortColumn-1], nil
	}
	return "", fmt.Errorf("can not find column")
}

func sortRows(rows []string, s *SortConfig) (string, error) {
	var sourceRows []string
	if s.isAlreadySorted {
		sourceRows = make([]string, len(rows))
		_ = copy(sourceRows, rows)

		// принудительно игнорируем все остальные флаги, если есть флаг -c
		s.sortColumn = 0
		s.sortByNumericValue = false
		s.reverseSort = false
		s.uniqueRows = false
	}

	switch {
	case s.sortColumn > 0:
		{
			// только уникальные строки
			if s.uniqueRows {
				rows = uniqualizer(rows)
			}

			sort.SliceStable(rows, func(i, j int) bool {
				// выбираем значение в выбранной колонке
				ith, err := getColumnValue(rows[i], s)
				if err != nil {
					return false
				}
				jth, err := getColumnValue(rows[j], s)
				if err != nil {
					return false
				}

				// сортировка по числам
				if s.sortByNumericValue {
					if s.reverseSort {
						return !ithNumLessThanJth(ith, jth)
					}
					return ithNumLessThanJth(ith, jth)
				}

				if s.reverseSort {
					return ith < jth
				}
				return ith > jth
			})
		}
	case s.sortByNumericValue:
		{
			// уникализируем, если надо
			if s.uniqueRows {
				rows = uniqualizer(rows)
			}

			sort.SliceStable(rows, func(i, j int) bool {
				if s.reverseSort {
					return !ithNumLessThanJth(rows[i], rows[j])
				}
				return ithNumLessThanJth(rows[i], rows[j])
			})
		}
	case s.uniqueRows:
		{
			rows = uniqualizer(rows)
			sort.SliceStable(rows, func(i, j int) bool {
				if s.reverseSort {
					return rows[i] > rows[j]
				}
				return rows[i] < rows[j]
			})
		}
	default:
		sort.SliceStable(rows, func(i, j int) bool {
			if s.reverseSort {
				return rows[i] > rows[j]
			}
			return rows[i] < rows[j]
		})
	}

	if s.isAlreadySorted {
		for i, row := range rows {
			if row != sourceRows[i] {
				return "false", nil
			}
		}
		return "true", nil
	}

	var result strings.Builder
	lenRows := len(rows)
	for i, row := range rows {
		if i < lenRows-1 {
			_, _ = result.WriteString(row + "\n")
		} else {
			_, _ = result.WriteString(row)
		}
	}

	return result.String(), nil
}

func ithNumLessThanJth(strI, strJ string) bool {
	ith, _ := strconv.Atoi(strI)
	jth, _ := strconv.Atoi(strJ)
	return ith < jth
}
