package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Stat("test.txt")

	if os.IsNotExist(err) {
		file, err := os.Create("test.txt")
		if err != nil {
			panic(err)
		}
		defer file.Close()

		fmt.Println("Success create file, %s")
		return
	} else {
		fmt.Println(("File already exists"))
	}
}
