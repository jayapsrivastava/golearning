package main

import (
  "flag"
  "fmt"
  "os"
	"github.com/jayapsrivastava/golearning/fileop"

)

func main() {

  flag.Parse()

  filename := flag.Arg(0)

  if filename == "" {
			fmt.Println("Usage: go run checksum.go <file_to_read>")
			os.Exit(1)
	}

	//Calculate checksum
  sum, err := fileop.Checksum(filename)
  if err != nil {
			fmt.Println(err)
	}
  fmt.Printf("From func - checksum: %x\n", sum)

	//Compress the file
  bytes, err := fileop.Compress(filename)
  if err != nil {
			fmt.Println(err)
	}
	fmt.Printf("From func - compress: %d\n", len(bytes))
}
