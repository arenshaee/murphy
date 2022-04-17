package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func getTodaysLaw() (law string) {
	laws := readLawFile("tech_laws.file")
	lineToRead := getLawNumberToRead(len(laws))

	return laws[lineToRead]
}

func getLawNumberToRead(max int) (lineToRead int) {
	for runCount := 0; ; runCount++ {
		rand.Seed(time.Now().UnixNano())
		lineToRead = rand.Intn(rand.Intn(max-1+1) + 1)
		statFileSlice := readStatFile()
		
		if getPosition(statFileSlice, strconv.Itoa(lineToRead)) == -1 {
			break
		}

		if runCount == max {
			clearStatFile()
			runCount = 0
		}
	}
	writeStatFile(fmt.Sprintln(lineToRead))

	return
}

func clearStatFile() {
	_, err := os.Create("stat")
	if err != nil {
		panic(err)
	}
}

func writeStatFile(input string) {
	// []byte(input), 0666
	f, err := os.OpenFile("stat", os.O_CREATE|os.O_APPEND, os.ModeAppend)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, werr := f.WriteString(input)
	if werr != nil {
		panic(werr)
	}
}

func readStatFile() (dataSlice []string) {
	f, err := os.Open("stat")

	if err != nil {
		fmt.Println("Ridi")
		return nil
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		dataSlice = append(dataSlice, scanner.Text())
	}

	return
}

func readLawFile(fileName string) (dataMap map[int]string) {
	f, err := os.Open(fileName)

	if err != nil {
		fmt.Println("Ridi")
		return nil
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	dataMap = make(map[int]string)
	lineNumber := 0
	for scanner.Scan() {
		lineNumber++
		dataMap[lineNumber] = scanner.Text()
	}

	return
}