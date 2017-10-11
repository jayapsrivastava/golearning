package main

import (
		"fmt"
		"os"
)

func main() {
		err := os.Remove("test.txt")
		if err != nil {
				fmt.Println(err)
		}
}

