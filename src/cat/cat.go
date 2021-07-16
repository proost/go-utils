package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	files := os.Args[1:]
	for _, file := range files {
		readFile(file)
	}
}

func readFile(filePath string) {
	data, err := os.Open(filePath)
	isPossibleToOpen(err)
	defer data.Close()

	fmt.Printf("File name: %s \n", filePath)
	_readFile(data)
}

func isPossibleToOpen(err error) {
	if err != nil {
		panic("cat: cannot open file")
	}
}

func _readFile(data *os.File) {
	scanner := bufio.NewScanner(data)

	for scanner.Scan() {
		println(scanner.Text())
	}
}