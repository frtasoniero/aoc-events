package daythree

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func Run() error {
	file, err := os.Open("./daythree/input.txt")

	if err != nil {
		fmt.Println("Error opening file:", err)
		return err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	regex := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	// var total float64 = 0.0

	for scanner.Scan() {
		line := scanner.Text()

		matches := regex.FindAllString(line, -1)
		for _, match := range matches {
			fmt.Println("Found match:", match)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return err
	}

	return nil
}
