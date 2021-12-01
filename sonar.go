package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func sonar() {
	file, err := os.Open("sonar_input.txt")

	if err != nil {
		log.Fatalf("Failed to open file!")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var prevDepth = 10000000000
	var timesIncreased = 0

	for scanner.Scan() {
		convertedInt, err := strconv.Atoi(scanner.Text())
		if err == nil {
			if convertedInt > prevDepth {
				timesIncreased++
			}
			prevDepth = convertedInt
		}
	}

	file.Close()

	log.Printf("%d", timesIncreased)
}

func main() {
	sonar()
}
