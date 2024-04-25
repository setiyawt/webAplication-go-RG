package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.Create("file.txt")
	if err != nil {
		log.Fatalf("error %v\n", err)
	}
}
