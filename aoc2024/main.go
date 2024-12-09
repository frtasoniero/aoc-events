package main

import (
	"fmt"
	"log"

	"github.com/frtasoniero/aoc-events/aoc2024/daytwo"
)

func main() {
	if err := daytwo.Run(); err != nil {
		log.Fatalf("Error running dayone: %v", err)
	}

	fmt.Println("Day one completed successfully")
}
