package daythree

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func Run() error {
	file, err := os.Open("./daythree/input.txt")

	if err != nil {
		fmt.Println("Error opening file:", err)
		return err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	regexWords := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	regexNumbers := regexp.MustCompile("[0-9]+")
	var multiplicationsResult float64 = 0.0

	for scanner.Scan() {
		line := scanner.Text()

		matches := regexWords.FindAllString(line, -1)
		for _, match := range matches {
			numbers := regexNumbers.FindAllString(match, -1)

			num1, err := strconv.Atoi(numbers[0])
			if err != nil {
				return fmt.Errorf("error converting number: %w", err)
			}
			num2, err := strconv.Atoi(numbers[1])
			if err != nil {
				return fmt.Errorf("error converting number: %w", err)
			}

			multiplicationsResult = multiplicationsResult + (float64(num1) * float64(num2))
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return err
	}

	fmt.Println("Multiplications result:", multiplicationsResult)

	return nil
}
