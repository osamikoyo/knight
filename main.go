package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("number of arguments so small, Usage: knight [pipe_name]")
		return
	}

	pipeName := os.Args[1]

	manifest, err := ParseFile()
	if err != nil {
		fmt.Printf("error parse manifest: %v\n", err)
		return
	}

	if err = manifest.Run(pipeName); err != nil {
		fmt.Printf("error run pipeline: %v", err)
		return
	}
}
