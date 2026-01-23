package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	root := "."
	if len(os.Args) > 1 {
		root = os.Args[1]
	}

	// Use CPU count * 2 for optimal I/O throughput
	engine := NewEngine(root, runtime.NumCPU()*2)
	results := make(chan string, 100)

	go engine.Walk(root, results)

	for path := range results {
		fmt.Println(path)
	}
}
