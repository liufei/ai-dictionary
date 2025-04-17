package main

import (
	"fmt"

	"github.com/liufei/ai-dictionary/cmd"
)

func main() {
	if _, err := cmd.NewCLI().Run(); err != nil {
		fmt.Printf("Error running program: %v\n", err)
	}
}
