package main

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

func TestRunCLI(t *testing.T) {
	inputData := "2+2\nexit\n"
	expectedOutput := "Введите арифметическое выражение (или 'exit' для выхода): Результат: 4.000000\n" +
		"Введите арифметическое выражение (или 'exit' для выхода): Выход из программы.\n"

	inputReader := bufio.NewReader(strings.NewReader(inputData))
	var outputBuffer bytes.Buffer
	outputWriter := bufio.NewWriter(&outputBuffer)

	runCLI(inputReader, outputWriter)

	outputWriter.Flush()

	actualOutput := outputBuffer.String()
	if actualOutput != expectedOutput {
		t.Errorf("Ожидаемый вывод:\n%s\nФактический вывод:\n%s", expectedOutput, actualOutput)
	}
}
