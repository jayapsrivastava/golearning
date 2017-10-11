package main

import (
  "crypto/md5"
  "log"
  "fmt"
  "io"
  "os"
)

type MyReader struct {
		r file
		i io.Reader
}

func (my *MyReader) Read(p []byte) (n int, err error) {

func main() {
  file, err := os.Open("test.txt")
  if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	
	hash := md5.New()
  nBytes, err := io.Copy(hash,file)
  if err != nil {
		log.Fatal(err)
	}
  fmt.Printf("No.of bytes: %d\n", nBytes)

	sum := hash.Sum(nil)
  fmt.Printf("Md5 Checksum: %x\n", sum)
}
	
