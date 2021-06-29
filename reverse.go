package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
)

func main() {
	inputFilePath, outputFilePath := parseInput()

	if inputFilePath == nil {
		return
	}
	inputFile, err := os.Open(*inputFilePath)
	if err != nil {
		panic(fmt.Sprintf("error: cannot open file %s", inputFile.Name()))
	}
	defer inputFile.Close()

	var outputFile *os.File
	if outputFilePath != nil {
		outputFile, err = os.OpenFile(*outputFilePath, os.O_CREATE|os.O_RDWR, 0644)
		if err != nil {
			panic(fmt.Sprintf("error: cannot open file %s", outputFile.Name()))
		}
	}
	defer outputFile.Close()

	reversedLines := reverseLines(inputFile)
	if outputFile == nil {
		for _, line := range reversedLines {
			println(line)
		}
	} else {
		writeFile(outputFile, reversedLines)
	}
}

func parseInput() (*string, *string) {
	args := os.Args[1:]

	if len(args) < 1 {
		return nil, nil
	} else if len(args) == 1 {
		return &args[0], nil
	} else if len(args) == 2 {
		return &args[0], &args[1]
	} else {
		panic("usage: reverse <input> <output>")
	}
}

func reverseLines(file *os.File) []string {
	scanner := bufio.NewScanner(file)

	stack := list.New()
	for scanner.Scan() {
		line := scanner.Text()
		stack.PushBack(line)
	}

	reversedStrings := make([]string, 0)
	for e := stack.Back(); e != nil; e = e.Prev() {
		line := e.Value.(string)
		reversedStrings = append(reversedStrings, line)
	}

	return reversedStrings
}

func writeFile(file *os.File, lines []string) {
	writer := bufio.NewWriter(file)

	for _, line := range lines {
		_ , err := writer.WriteString(line + "\n")
		if err != nil {
			panic(err)
		}
	}

	writer.Flush()
}
