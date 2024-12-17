package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/MoonLeeg/calc-app/internal/calculator"
	"github.com/MoonLeeg/calc-app/internal/handler"
)

func main() {
	go startWebServer()
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	runCLI(reader, writer)
}

func runCLI(inputReader *bufio.Reader, outputWriter *bufio.Writer) {
	for {
		fmt.Fprint(outputWriter, "Введите арифметическое выражение (или 'exit' для выхода): ")
		outputWriter.Flush()

		input, err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(outputWriter, "Ошибка чтения ввода:", err)
			outputWriter.Flush()
			continue
		}

		input = strings.TrimSpace(input)
		if input == "exit" {
			fmt.Fprintln(outputWriter, "Выход из программы.")
			outputWriter.Flush()
			break
		}

		result, err := calculator.Calc(input)
		if err != nil {
			fmt.Fprintln(outputWriter, "Ошибка вычисления:", err)
			outputWriter.Flush()
			continue
		}

		fmt.Fprintf(outputWriter, "Результат: %f\n", result)
		outputWriter.Flush()
	}
}

func startWebServer() {
	http.HandleFunc("/api/v1/calculate", handler.CalculateHandler)
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("ListenAndServe: %v", err)
	}
}
