package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func moveSubmarine(x *int, y *int, value string) {
	splitValue := strings.Fields(value)
	moveAmount, _ := strconv.Atoi(splitValue[1])

	switch splitValue[0] {
	case "up":
		*y = *y - moveAmount
	case "down":
		*y = *y + moveAmount
	case "forward":
		*x = *x + moveAmount
	}
}

func movement() {
	x := 0
	y := 0

	file, err := os.Open("movement_input.txt")

	if err != nil {
		log.Fatalf("Failed to open file!")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		moveSubmarine(&x, &y, scanner.Text())
	}

	log.Printf("%d", x*y)
}

func main() {
	movement()
}
