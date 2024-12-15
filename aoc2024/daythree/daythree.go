package daythree

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Run() error {
	file, err := os.Open("./daythree/input.txt")

	if err != nil {
		fmt.Println("Error opening file:", err)
		return err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var fileInformation strings.Builder

	var multiplicationsResult int64 = 0.0

	for scanner.Scan() {
		line := scanner.Text()

		fileInformation.WriteString(line)
	}

	doSeparator := "do()"
	dontSeparator := "don't()"

	splitByDo := strings.Split(fileInformation.String(), doSeparator)

	for _, do := range splitByDo {
		splitByDont := strings.Split(do, dontSeparator)

		result, err := computeMulFromString(splitByDont[0])

		if err != nil {
			fmt.Println("Error computing mul:", err)
			return err
		}
		multiplicationsResult += result
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return err
	}

	fmt.Println("Multiplications result:", multiplicationsResult)

	return nil
}

func computeMulFromString(line string) (int64, error) {
	regexMulWord := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	regexNumbers := regexp.MustCompile("[0-9]+")

	matchesMul := regexMulWord.FindAllString(line, -1)

	var result int64 = 0.0

	for _, match := range matchesMul {

		numbers := regexNumbers.FindAllString(match, -1)

		num1, err := strconv.Atoi(numbers[0])
		if err != nil {
			return 0, fmt.Errorf("error converting number: %w", err)
		}
		num2, err := strconv.Atoi(numbers[1])
		if err != nil {
			return 0, fmt.Errorf("error converting number: %w", err)
		}

		result += (int64(num1) * int64(num2))
	}

	return result, nil
}
