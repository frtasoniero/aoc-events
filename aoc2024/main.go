package main

import (
	"fmt"
	"log"

	"github.com/frtasoniero/aoc-events/aoc2024/dayone"
)

func main() {
	if err := dayone.Run(); err != nil {
		log.Fatalf("Error running dayone: %v", err)
	}

	fmt.Println("Day one completed successfully")
}
