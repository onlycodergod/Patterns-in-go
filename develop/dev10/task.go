package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

/*
=== Утилита telnet ===
Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123
Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).
При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.

Пример использования:
go run task.go --timeout=10s opennet.ru  80
GET /

*/

func main() {
	timeout := flag.Duration("timeout", 10*time.Second, "timeout")
	flag.Parse() // флаги из консольной строки

	args := flag.Args() // аргументы из консольной строке

	if len(args) != 2 {
		log.Fatalf("Enter: go-telnet --timeout=10s host port")
	}

	host := flag.Arg(0)
	port := flag.Arg(1)

	conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), *timeout)
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}
	defer conn.Close()

	for {
		err = Write(conn)
		if err != nil {
			if err == io.EOF {
				fmt.Println("Finish")
				os.Exit(0)
			}
			return
		}

		err = Read(conn)
		if err != nil {
			fmt.Println("Server close connection")
			return
		}
	}

}

func Write(conn net.Conn) error {
	// читаем из stdin, записываем в conn
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Text to send: ")
	text, err := reader.ReadString('\n')
	if err != nil {
		return err
	}

	_, err = conn.Write([]byte(text))
	if err != nil {
		return err
	}

	return nil
}

func Read(conn net.Conn) error {
	// читаем из conn, записываем в stdout
	reader := bufio.NewReader(conn)
	text, err := reader.ReadString('\n')
	if err != nil {
		return err
	}

	_, err = fmt.Print(text)
	if err != nil {
		return err
	}

	return nil
}
