package dayone

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Run() error {
	file, err := os.Open("./dayone/input.txt")

	if err != nil {
		fmt.Println("Error opening file:", err)
		return err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	columnOne := make([]int, 0, 1000)
	columnTwo := make([]int, 0, 1000)
	columnTwoCount := make(map[int]int)

	for scanner.Scan() {
		line := scanner.Text()

		numbers := strings.Fields(line)

		num1, err := strconv.Atoi(numbers[0])
		if err != nil {
			return fmt.Errorf("error converting number: %w", err)
		}
		num2, err := strconv.Atoi(numbers[1])
		if err != nil {
			return fmt.Errorf("error converting number: %w", err)
		}

		columnOne = append(columnOne, num1)
		columnTwo = append(columnTwo, num2)
		columnTwoCount[num2]++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return err
	}

	sort.Ints(columnOne)
	sort.Ints(columnTwo)

	distanceResult := 0.0

	for i := range len(columnOne) {
		distanceResult += math.Abs(float64(columnOne[i] - columnTwo[i]))
	}

	fmt.Println("Result of distances:", distanceResult)

	similarityResult := 0

	for _, num := range columnOne {
		similarityResult += num * columnTwoCount[num]
	}

	fmt.Println("Result of similarity:", similarityResult)

	return nil
}
