package main

import (
	"fmt"
	"os"
	"terminal-ai-agent/internal/ui"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Warning: Could not load .env file")
	}

	p := ui.NewProgram()
	if err := p.Start(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
