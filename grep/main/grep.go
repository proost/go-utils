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
		readFile(file, keyword)
	}
}

func readFile(filePath string, keyword string) {
	data, err := os.Open(filePath)
	isPossibleToOpen(err)
	defer data.Close()

	findKeywordContainedLine(data, keyword)
}

func isPossibleToOpen(err error) {
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