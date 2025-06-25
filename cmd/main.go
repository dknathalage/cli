package main

import (
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Println("Usage: cli <command>")
		os.Exit(1)
	}

	if os.Args[1] == "gen" {
		log.Println("Generating config code")
	}

	if os.Args[1] == "init" {
		log.Println("Initializing config code")
	}

	if os.Args[1] == "apply" {
		log.Println("Applying config code")
	}
}
