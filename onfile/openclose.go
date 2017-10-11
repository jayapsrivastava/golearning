package main

import (
		"fmt"
		"os"
)

func main() {
		file, err := os.OpenFile("openclose.txt", os.O_RDONLY| os.O_CREATE, 0666)
    if err != nil {
				fmt.Println(err)
		}
		file.Close()
}

