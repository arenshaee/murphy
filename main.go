package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	getTodaysLaw()
}

func getTodaysLaw() (law string) {
	lastReadLine := getLastReadLineNumber()
	// laws := readFile("tech_laws.file")
	fmt.Println(lastReadLine)

	return ""
}

func getLastReadLineNumber() int {
	const LAST_READ_LINE_ID = 1

	statFileOutput := readFile("stat")
	
	lastReadLine, err := strconv.Atoi(statFileOutput[LAST_READ_LINE_ID])

	if err != nil {
		fmt.Println(err)
		lastReadLine = 1
		writeFile(strconv.Itoa(lastReadLine))
	}

	return lastReadLine
}

func writeFile(input string) {
	err := os.WriteFile("stat", []byte(input), 0666)
	if err != nil {
		panic(err)
	}
}

func readFile(fileName string) map[int]string {
	f, err := os.Open(fileName)

	if err != nil {
		fmt.Println("Ridi")
		return nil
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	dataMap := make(map[int]string)
	lineNumber := 0
	for scanner.Scan() {
		lineNumber++
		dataMap[lineNumber] = scanner.Text()
	}

	return dataMap
}
