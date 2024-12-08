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

func Run(path string) {
	file, err := os.Open(path)

	if err != nil {
		fmt.Println("Error opening file:", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineCount := 0
	columnOne := make([]int, 0, 1000)
	columnTwo := make([]int, 0, 1000)
	columnTwoCount := make(map[int]int)

	for scanner.Scan() {
		line := scanner.Text()
		lineCount++

		numbers := strings.Fields(line)
		columnOne = append(columnOne, convertToInt(numbers[0]))
		columnTwo = append(columnTwo, convertToInt(numbers[1]))
		columnTwoCount[convertToInt(numbers[1])]++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	sort.Ints(columnOne)
	sort.Ints(columnTwo)

	distanceResult := 0.0

	for i := range len(columnOne) {
		distanceResult += math.Abs(float64(columnOne[i] - columnTwo[i]))
	}

	fmt.Println("Result of distances:", distanceResult)

	similarityResult := 0

	for i := range len(columnOne) {
		similarityResult += columnOne[i] * columnTwoCount[columnOne[i]]
	}

	fmt.Println("Result of similarity:", similarityResult)
}

func convertToInt(number string) int {
	numberStr := number
	numberInt, err := strconv.Atoi(numberStr)

	if err != nil {
		fmt.Println("Error during conversion")
		return -1
	}

	return numberInt
}
