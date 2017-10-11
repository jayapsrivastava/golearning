package fileop

import (
  "compress/gzip"
  "crypto/md5"
  "io/ioutil"
  "bufio"
  "bytes"
  "fmt"
  "io"
  "os"
)

func Checksum(filename string) ([]byte, error) {
  file, err := os.Open(filename)
  if err != nil {
    fmt.Println(err)
  }
  defer file.Close()

  hash := md5.New()
  _, err := io.Copy(hash,file)
  if err != nil {
    return nil, io.ErrUnexpectedEOF
  }

  sum := hash.Sum(nil)
  fmt.Printf("Md5 Checksum: %x\n", sum)
  return sum, nil
}

func Compress(filename string) ([]byte, error) {
  file, err := os.Open(filename)
  if err != nil {
    fmt.Println(err)
  }
  defer file.Close()

  fileInfo, err := file.Stat()
  size := fileInfo.Size()
  gzf := filename+".gz"

  //Read file data into buffer
  nBytes := make([]byte, size)
  buffer := bufio.NewReader(file)
  _, err = buffer.Read(nBytes)
  if err != nil {
    fmt.Printf("Failed to read file :%s into buffer and error: %s\n", filename, err)
    return nil, err
  }

  //Create a gzip writer on top of file writer
  var buf bytes.Buffer
  gzw := gzip.NewWriter(&buf)
  gzw.Write(nBytes)
  gzw.Close()

  err = ioutil.WriteFile(gzf, buf.Bytes(), 06444)
  if err != nil {
    fmt.Printf("Failed to write to file :%s and error: %s\n", gzf, err)
    return nil, err
  }
  fmt.Printf("Compressed file: %s with bytes:%d\n", gzf, buf.Bytes)
  return buf.Bytes(), nil
}

