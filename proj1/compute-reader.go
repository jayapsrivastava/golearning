package main

import (
  "path/filepath"
  "compress/gzip"
  "crypto/cipher"
  "crypto/aes"
  "flag"
  "fmt"
  "io"
  "os"
)

func main() {
  flag.Parse()
  fgz := flag.Arg(0)

  if fgz == "" {
    fmt.Println("Usage: go run compute-reader.go <file_to_read>")
    os.Exit(1)
}

  //Open gzip file
  fin, err := os.Open(fgz)
  if err != nil {
    fmt.Println(err)
  }
  defer fin.Close()

  zr, err := gzip.NewReader(fin)
  if err != nil {
    fmt.Println(err)
  }
  defer zr.Close()

  //gzip.Header struct consist of metadata information, zr.Name = file name
  var file string
  file = filepath.Join(file, zr.Name)
  fout, err := os.Create(file)
  if err != nil {
    fmt.Println(err)
  }
  defer fout.Close()

  //md5 can be used for key but for simplicity using this
  key := make([]byte, 16)
  cblock, err := aes.NewCipher(key)
  if err != nil {
    fmt.Println(err)
  }

  var iv [aes.BlockSize]byte
  stream := cipher.NewOFB(cblock, iv[:])

  reader := cipher.StreamReader{S: stream, R: zr}
  _, err = io.Copy(fout, reader)
  if err != nil {
    fmt.Println(err)
  }

}
