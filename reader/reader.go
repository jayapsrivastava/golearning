package main

import "golang.org/x/tour/reader"
        
type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

func (m MyReader) Read(b []byte) (int, error){
	if nil == b {
		return 0, nil
	}
	cnt := 0
	for i, _ := range b{
		b[i] = 'A'
		cnt = cnt + 1
	}
	return cnt, nil
}

func main() {
	reader.Validate(MyReader{})
}
