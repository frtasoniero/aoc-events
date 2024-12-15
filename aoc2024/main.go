package main

import (
	"fmt"
	"log"

	"github.com/frtasoniero/aoc-events/aoc2024/daythree"
)

func main() {
	if err := daythree.Run(); err != nil {
		log.Fatalf("Error running dayone: %v", err)
	}

	fmt.Println("Day three completed successfully")
}
