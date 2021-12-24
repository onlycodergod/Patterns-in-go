package main

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===
Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.
Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	words := []string{"Пятак", "тяпка", "пятка", "листок", "тяпка", "Слиток", "столик", "приказ", "каприз", "одно"}
	fmt.Println(SearchAnagrams(words))
}

func SearchAnagrams(data []string) map[string][]string {
	temp := make(map[string][]string)
	res := make(map[string][]string)
	for _, str := range data {
		lowStr := strings.ToLower(str)       // приведение к нижнему регистру
		sortStr := SortLettersInWord(lowStr) // сортировка всех букв в слове
		if _, ok := temp[sortStr]; ok {      // поиск ключа отсортированного слова во временном множестве
			isUnique := true
			for _, val := range temp[sortStr] {
				if val == lowStr { // проверяем на уникальность значения во множестве
					isUnique = false
				}
			}
			if isUnique {
				temp[sortStr] = append(temp[sortStr], lowStr) // добавляем уникальное обычное слово в нижнем регистре во множество с нужным ключом
			}
		} else {
			temp[sortStr] = []string{lowStr} // добавляем во временное множество (ключ => значение): отсортированное слово => обычное слово в нижнем регистре
		}
	}
	for _, v := range temp {
		if len(v) <= 1 {
			continue
		}
		res[v[0]] = v[1:]
	}
	return res
}

func SortLettersInWord(str string) string {
	r := []rune(str)
	sort.Sort(runes(r))
	return string(r)
}

type runes []rune

func (r runes) Len() int {
	return len(r)
}
func (r runes) Swap(i, j int) {
	swap := reflect.Swapper(r)
	swap(i, j)
}
func (r runes) Less(i, j int) bool {
	return r[i] < r[j]
}
