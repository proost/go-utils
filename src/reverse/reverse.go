package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	lines := parseFile(inputFile)
	sort.Sort(sort.Reverse(sort.StringSlice(lines)))
	if outputFile == nil {
		for _, line := range lines {
			println(line)
		}
	} else {
		writeFile(outputFile, lines)
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

func parseFile(file *os.File) []string {
	result := make([]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	return result
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
