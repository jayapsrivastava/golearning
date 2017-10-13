package main

import (
  "path/filepath"
  "compress/gzip"
  "crypto/cipher"
  "crypto/aes"
  "crypto/md5"
  "flag"
  "fmt"
  "io"
  "os"
)

func main() {
  flag.Parse()
  filename := flag.Arg(0)

  if filename == "" {
    fmt.Println("Usage: go run compute-writer.go <file_to_compute>")
    os.Exit(1)
  }

  //Open a file for reading
  r, err := os.Open(filename)
  if err != nil {
    fmt.Println(err)
  }
  defer r.Close()

  //Calculate the checksum
  //r := strings.NewReader("Calculate the checksum")
  hash := md5.New()
  _, err = io.Copy(hash, r)
  if err != nil {
    fmt.Println(err)
  }
  sum := hash.Sum(nil)
  fmt.Printf("md5 Checksum: %x\n", sum)

  //Open a file for again reading
  r, err = os.Open(filename)
  if err != nil {
    fmt.Println(err)
  }
  defer r.Close()

  //Create a .gz file to write
  var fgz string
  fgz = filepath.Join(fgz,fmt.Sprintf("%s.gz", filename))
  f, err := os.Create(fgz)
  if err != nil {
    fmt.Println(err)
  }
  defer f.Close()

  //Create a gzip writer gzip.Header struct consist of 
  //metadata information, zr.Name = file name
  zw := gzip.NewWriter(f)
  zw.Name = filename
  defer zw.Close()

  // For encrypt  md5 can be used for key but for simplicity using this
  //  key := sum
  key := make([]byte, 16)
  cblock, err := aes.NewCipher(key)
  if err != nil {
    fmt.Println(err)
  }

  var iv [aes.BlockSize]byte
  stream := cipher.NewOFB(cblock, iv[:])
  writer := cipher.StreamWriter{S: stream, W: zw}
  _, err = io.Copy(writer, r)
  if err != nil {
    fmt.Println(err)
  }
}
