package main

import "golang.org/x/tour/reader"

type MyReader struct{}

func (reader MyReader) Read(streams []byte) (int, error) {
	for index := range streams {
		streams[index] = 'A'	
	}
	
	return len(streams), nil
}
// TODO: Add a Read([]byte) (int, error) method to MyReader.

func main() {
	reader.Validate(MyReader{})
}
