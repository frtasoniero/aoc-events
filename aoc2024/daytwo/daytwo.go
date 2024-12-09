package daytwo

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func Run() error {
	file, err := os.Open("./daytwo/input.txt")

	if err != nil {
		return fmt.Errorf("error opening file: %w", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	safetyReport := make(map[bool]int)

	for scanner.Scan() {
		line := scanner.Text()

		isIncreasing := false
		isDecreasing := false
		isSafe := true

		numbers := strings.Fields(line)

		for i := 0; i < len(numbers)-1; i++ {
			num1, err := strconv.Atoi(numbers[i])
			if err != nil {
				return fmt.Errorf("error converting number: %w", err)
			}

			num2, err := strconv.Atoi(numbers[i+1])
			if err != nil {
				return fmt.Errorf("error converting number: %w", err)
			}

			if num1 < num2 {
				isIncreasing = true
			} else if num1 > num2 {
				isDecreasing = true
			}

			if num1 == num2 || math.Abs(float64(num1-num2)) > 3 || isIncreasing == isDecreasing {
				isSafe = false
				safetyReport[false]++
				break
			}
		}

		if isSafe {
			safetyReport[true]++
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}

	fmt.Println("Safety Report:", safetyReport)

	return nil
}
