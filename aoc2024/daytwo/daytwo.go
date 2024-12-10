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

		isSafe := true

		numbers := strings.Fields(line)

		isSafe, err = checkSafety(numbers, isSafe)

		if err != nil {
			return fmt.Errorf("error checking safety: %w", err)
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

func checkSafety(numbers []string, isSafe bool) (bool, error) {
	isIncreasing := false
	isDecreasing := false

	for i := 0; i < len(numbers)-1; i++ {
		num1, err := strconv.Atoi(numbers[i])
		if err != nil {
			return false, err
		}

		num2, err := strconv.Atoi(numbers[i+1])
		if err != nil {
			return false, err
		}

		if num1 < num2 {
			isIncreasing = true
		} else if num1 > num2 {
			isDecreasing = true
		}

		if num1 == num2 || math.Abs(float64(num1-num2)) > 3 {
			if isSafe {
				numbersCopy := make([]string, len(numbers))
				copy(numbersCopy, numbers)

				numbers1 := append(numbersCopy[:i], numbersCopy[i+1:]...)
				check1, err1 := checkSafety(numbers1, false)

				if err1 != nil {
					return false, err1
				}

				numbersCopy = make([]string, len(numbers))
				copy(numbersCopy, numbers)

				numbers2 := append(numbersCopy[:i+1], numbersCopy[i+2:]...)
				check2, err2 := checkSafety(numbers2, false)

				if err2 != nil {
					return false, err2
				}

				if check1 || check2 {
					return true, nil
				} else {
					return false, nil
				}
			}

			return false, nil
		} else if isIncreasing == isDecreasing {
			if isSafe {
				numbersCopy := make([]string, len(numbers))
				copy(numbersCopy, numbers)

				numbers0 := append(numbersCopy[:i-1], numbersCopy[i:]...)
				check0, err0 := checkSafety(numbers0, false)

				if err0 != nil {
					return false, err0
				}

				numbersCopy = make([]string, len(numbers))
				copy(numbersCopy, numbers)

				numbers1 := append(numbersCopy[:i], numbersCopy[i+1:]...)
				check1, err1 := checkSafety(numbers1, false)

				if err1 != nil {
					return false, err1
				}

				numbersCopy = make([]string, len(numbers))
				copy(numbersCopy, numbers)

				numbers2 := append(numbersCopy[:i+1], numbersCopy[i+2:]...)
				check2, err2 := checkSafety(numbers2, false)

				if err2 != nil {
					return false, err2
				}

				if check0 || check1 || check2 {
					return true, nil
				} else {
					return false, nil
				}
			}

			return false, nil
		}
	}

	return true, nil
}
