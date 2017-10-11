package main

import (
//		"fmt"
		"os"
    "log"
)

func main() {
		oldpath := "test.txt"
		newpath := "rename.txt"

		err := os.Rename(oldpath, newpath)
		if err != nil {
				//fmt.Println(err)
        log.Fatal(err)
		}
		
}

