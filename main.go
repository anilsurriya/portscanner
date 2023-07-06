package main

import (
	"fmt"

	"github.com/anilsurriya/portscanner/scanner"
)

func main() {

	sj := scanner.NewJob("tcp", "127.0.0.1", 18000, 20000, 3)

	out := sj.StartJob()

	openCount := 0

	for _, result := range out {
		fmt.Println(&result)
		if result.Status == scanner.Open {
			openCount++
		}
	}

	fmt.Printf("Total open ports: %v", openCount)
}
