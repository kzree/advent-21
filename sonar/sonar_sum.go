package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func sumArray(array []int) int {
	sum := 0
	for i := 0; i < len(array); i++ {
		sum += array[i]
	}

	return sum
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func sonarSum() {
	file, err := os.Open("sonar_input.txt")

	if err != nil {
		log.Fatalf("Failed to open file!")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lastSum int = 10000000
	var measurements []int
	var timesIncreased int = 0

	for scanner.Scan() {
		convertedInt, err := strconv.Atoi(scanner.Text())
		if err == nil {
			measurements = append(measurements, convertedInt)
			if len(measurements) == 3 {
				sum := sumArray(measurements)
				if sum > lastSum {
					timesIncreased++
				}
				lastSum = sum
				measurements = remove(measurements, 0)
			}
		}
	}

	log.Printf("%d", timesIncreased)
}

func main() {
	sonarSum()
}
