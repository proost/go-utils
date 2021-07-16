package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	keyword := os.Args[1]
	files := os.Args[2:]
	for _, file := range files {
		readFileWithKeyword(file, keyword)
	}
}

func readFileWithKeyword(filePath string, keyword string) {
	data, err := os.Open(filePath)
	isFilePossibleToOpen(err)
	defer data.Close()

	findKeywordContainedLine(data, keyword)
}

func isFilePossibleToOpen(err error) {
	if err != nil {
		panic("cat: cannot open file")
	}
}

func findKeywordContainedLine(data *os.File, keyword string) {
	scanner := bufio.NewScanner(data)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, keyword) {
			fmt.Println(line)
		}
	}
}