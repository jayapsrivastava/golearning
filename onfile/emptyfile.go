package main

import (
	"fmt"
	"os"
)

func main() {
	newFile, err := os.Create("test.txt")
	if err != nil {
    fmt.Println(err)
	}
	fmt.Println(newFile)
	fmt.Printf("\nFile name: %s\n", newFile.Name())
	newFile.Close()
}
