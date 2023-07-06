package main

import (
	"fmt"

	"github.com/anilsurriya/portscanner/scanner"
)

func main() {
	fmt.Println("...")

	sj := scanner.NewJob("tcp", "127.0.0.1", 18000, 20000, 3)

	sj.Generate()
	sj.StartWorkers()

	openCount := 0

	for result := range sj.Results {
		fmt.Println(&result)
		if result.Status == scanner.Open {
			openCount++
		}
	}

	fmt.Printf("Total open ports: %v", openCount)
}
