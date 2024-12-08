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

	for scanner.Scan() {
		line := scanner.Text()
		lineCount++

		numbers := strings.Fields(line)
		columnOne = append(columnOne, convertToInt(numbers[0]))
		columnTwo = append(columnTwo, convertToInt(numbers[1]))

		fmt.Println(line)

		// if lineCount >= 50 {
		// 	break
		// }
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	sort.Ints(columnOne)
	sort.Ints(columnTwo)

	result := 0.0

	for i := range len(columnOne) {
		result += math.Abs(float64(columnOne[i] - columnTwo[i]))
	}

	fmt.Println("Result:", result)
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
