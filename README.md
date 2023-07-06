# portscanner

## usage
import the package scanner from github.com/anilsurriya/portscanner/scanner

- NewJob returns a scanJob struct with the state parameters to start the scan.
- StartJob method on scanJob is used to start and wait for the completion of the scan job.

## code
```go
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
```

## working
- A generator piplelined with concurrent worker goroutines is used.
- The generator generates addresses to scan and one of the worker gorountine scans it concurrently.
- The results are piped into an unbuffered channel.

## To-Do
[] UI
