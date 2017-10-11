package main

import (
		"fmt"
		"os"
)


func main() {
		fileinfo, err := os.Stat("test.txt")
    if err != nil {
				fmt.Println(err)
		}
		fmt.Println("File name:\n", fileinfo.Name())
		fmt.Println("File size:\n", fileinfo.Size())
		fmt.Println("File Mode:\n", fileinfo.Mode())
		fmt.Println("File ModTime:\n", fileinfo.ModTime())
		fmt.Println("IsDir: \n", fileinfo.IsDir())
		fmt.Printf("System Interface info: %+v\n", fileinfo.Sys())
		fmt.Printf("System Interface:%T\n", fileinfo.Sys())
    fmt.Println(fileinfo)
}

