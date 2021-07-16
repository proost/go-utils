package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	filePaths := os.Args[1:]

	if len(filePaths) < 1 {
		panic("Usage: zip [read files]")
	}

	for _, filePath := range filePaths {
		process(filePath)
	}
}

func process(filePath string) {
	 file,err := os.Open(filePath)
	 if err != nil {
	 	panic( fmt.Sprintf("Can't open file %s", filePath))
	 }
	 defer file.Close()

	 scanner := bufio.NewScanner(file)

	 for scanner.Scan() {
	 	line := compressFile(scanner.Text())
	 	fmt.Println(line)
	 }
}

func compressFile(line string) string {
	var curr rune = 0
	var result = ""
	var count = 0
	for _, char := range line {
		if curr == 0 {
			curr = char
			count = 1
			continue
		}

		if char != curr {
			result += strconv.Itoa(count) + string(curr)
			curr = char
			count = 1
		} else {
			count++
		}
	}

	result += strconv.Itoa(count) + string(curr)

	return result
}
