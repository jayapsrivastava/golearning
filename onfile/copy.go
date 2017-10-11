package main

import (
		"log"
		"os"
		"io"
)

func main() {
		orgFile, err := os.OpenFile("test.txt", os.O_RDONLY, 0666)
		if err != nil {
				log.Fatal(err)
		}
		defer orgFile.Close()

		newFile, err := os.Create("copy.txt")
		if err != nil {
				log.Fatal(err)
		}
    defer newFile.Close()
   
		nBytes, err := io.Copy(newFile, orgFile)
		if err != nil {
				log.Fatal(err)
		}
    log.Printf("Copied Bytes: %d\n", nBytes)
		
		err = newFile.Sync()
		if err != nil {
				log.Fatal(err)
		}
}
