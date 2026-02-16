package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const OvenTime = 40

func RemainingOvenTime(actual int) int {
	rem := OvenTime - actual
	if rem < 0 {
		return 0
	}
	return rem
}

func PreparationTime(numberOfLayers int) int {
	return numberOfLayers * 2
}

func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	return PreparationTime(numberOfLayers) + actualMinutesInOven
}

// ✅ Reusable input function
func readPositiveInt(reader *bufio.Reader, prompt string) int {
	for {
		fmt.Print(prompt)

		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		}

		line = strings.TrimSpace(line)
		value, err := strconv.Atoi(line)
		if err != nil || value < 0 {
			fmt.Println("Please enter a valid non-negative number.")
			continue
		}

		return value
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	layers := readPositiveInt(reader, "Enter the number of lasagna layers: ")
	minutes := readPositiveInt(reader, "Enter minutes already in the oven: ")

	fmt.Printf("Preparation time: %d minutes\n", PreparationTime(layers))
	fmt.Printf("Elapsed time: %d minutes\n", ElapsedTime(layers, minutes))
	fmt.Printf("Remaining oven time: %d minutes\n", RemainingOvenTime(minutes))
}
