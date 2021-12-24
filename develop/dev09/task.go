package main

import (
	"flag"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

/*
=== Утилита wget ===
Реализовать утилиту wget с возможностью скачивать сайты целиком
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.

// go run task.go https://ya.ru
*/

func main() {
	level := flag.Int("l", 0, "level")
	flag.Parse()
	url := flag.Arg(0)
	fileName := flag.Arg(1)

	err := wget(url, fileName, *level)
	if err != nil {
		log.Fatal(err)
	}
}

// запись запроса GET в файл
func wget(url string, fileName string, level int) (err error) {
	if level < 0 {
		return
	}

	// если для fileName задана пустая строка, то  последний фрагмент входного URL-адреса используется как имя файла
	if fileName == "" {
		urlSplit := strings.Split(url, "/")
		fileName = urlSplit[len(urlSplit)-1]
	}

	content, err := getContent(url)
	if err != nil {
		return
	}

	links := getLinks(content)

	err = writeToFile(fileName, content)
	if err != nil {
		return
	}

	for _, link := range links {
		err = wget(link, "", level-1)
		if err != nil {
			return
		}
	}
	return
}

//GET запрос на url - адрес
func getContent(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return content, nil
}

// Скачать и сохранить сайт в файл
func writeToFile(fileName string, content []byte) error {
	file, err := os.Create(fileName) // создать новый файл
	if err != nil {
		return err
	}
	defer file.Close() // закрыть файл

	_, err = file.Write(content)
	if err != nil {
		return err
	}
	return nil
}

// Собрать ссылки на странице
func getLinks(content []byte) (urls []string) {
	linkPattern := regexp.MustCompile(`(http|https):([\w\-_]+(?:(?:\.[\w\-_]+)+))([\w\-,@?^=%&amp;:/~#]*[\w\-?^=%&amp;/~#])?`)
	result := linkPattern.FindAll(content, -1)

	for _, url := range result {
		urls = append(urls, string(url))
	}
	return
}
