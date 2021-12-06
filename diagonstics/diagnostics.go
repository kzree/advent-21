package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Count struct {
	oneCount  int
	zeroCount int
}

func readFile() []string {
	var lines []string
	file, err := os.Open("diagnostics_input.txt")

	if err != nil {
		log.Fatalf("Failed to open file!")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func initCounter(length int) []Count {
	var counter []Count
	var initedItem Count
	initedItem.oneCount = 0
	initedItem.zeroCount = 0

	for i := 0; i < length; i++ {
		counter = append(counter, initedItem)
	}

	return counter
}

func splitCode(code string) []string {
	splitString := strings.Split(code, "")
	return splitString
}

func buildBinaryNumbers(counter []Count) [2]string {
	var gamma strings.Builder
	var epsilon strings.Builder

	for _, count := range counter {
		if count.oneCount > count.zeroCount {
			gamma.WriteString("1")
			epsilon.WriteString("0")
		} else {
			gamma.WriteString("0")
			epsilon.WriteString("1")
		}
	}

	answerArray := [2]string{gamma.String(), epsilon.String()}
	return answerArray
}

func getAnswer(values [2]string) {
	gamma, _ := strconv.ParseInt(values[0], 2, 64)
	epsilon, _ := strconv.ParseInt(values[1], 2, 64)

	log.Println(gamma * epsilon)
}

func main() {
	const codeLength = 12
	codes := readFile()
	counter := initCounter(codeLength)

	for _, code := range codes {
		codeArray := splitCode(code)

		for idx, number := range codeArray {
			convertedInt, _ := strconv.Atoi(number)
			if convertedInt == 0 {
				counter[idx].zeroCount++
			} else {
				counter[idx].oneCount++
			}
		}
	}

	getAnswer(buildBinaryNumbers(counter))
}
