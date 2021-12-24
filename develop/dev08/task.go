package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

/*
=== Взаимодействие с ОС ===
Необходимо реализовать собственный шелл
встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах
Реализовать утилиту netcat (nc) клиент
принимать данные из stdin и отправлять в соединение (tcp/udp)
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	reader := bufio.NewReader(os.Stdin) // читаем из stdin
	for {
		fmt.Print("$ ")
		cmdString, err := reader.ReadString('\n') // строка из консоли, разделение - перенос на новую строку
		if err != nil {
			_, err = fmt.Fprintln(os.Stderr, err)
			if err != nil {
				return
			}
		}

		err = runCommand(cmdString) // выполнение команды
		if err != nil {
			_, err = fmt.Fprintln(os.Stderr, err)
			if err != nil {
				return
			}
		}
	}
}

func runCommand(commandStr string) error {
	commandStr = strings.TrimSpace(commandStr)        // удаление пробелов в начале и в конце
	commandStr = strings.TrimSuffix(commandStr, "\n") // удаление символа переноса строки \n из строки консоли
	arrCommandStr := strings.Fields(commandStr)       // разделение на подстроки (отдельные слова)

	if len(arrCommandStr) == 0 { // если из слов ничего не осталось - возвращаем nil
		return nil
	}

	switch arrCommandStr[0] {
	// обработка перехода по директориям
	case "cd":
		if len(arrCommandStr) < 2 {
			return nil
		}

		err := os.Chdir(arrCommandStr[1])
		if err != nil {
			return err
		}
		return nil
	// обработка выхода из "шелла"
	case "\\quit":
		os.Exit(0)
	}

	cmd := exec.Command("bash", "-c", commandStr) // exec запускает выполнение команд из commandStr
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
