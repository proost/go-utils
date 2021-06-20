package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {
	filePath := os.Args[1]

	startProcess(filePath)
}

func startProcess(filePath string) {
	file,err := os.Open(filePath)
	if err != nil {
		panic( fmt.Sprintf("Can't open file %s", filePath))
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := decompressFile(scanner.Text())
		fmt.Println(line)
	}
}

func decompressFile(line string) string {
	var length = ""
	var result = ""
	for _, char := range line {
		if unicode.IsDigit(char) {
			length += string(char)
		}

		if unicode.IsLetter(char) {
			cnt, _ := strconv.Atoi(length)
			i := 0
			char := string(char)
			for i < cnt {
				result += char
				i++
			}

			length = ""
		}
	}

	return result
}
