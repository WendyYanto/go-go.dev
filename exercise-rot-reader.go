package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

const substitutionManipulator byte = 13

// references:
// key:   ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz
// value: NOPQRSTUVWXYZABCDEFGHIJKLMnopqrstuvwxyzabcdefghijklm

func (r rot13Reader) shouldAdd(value int32) bool {
	return ('A' <= value && value <= 'M') || ('a' <= value && value <= 'm')
}

func (r rot13Reader) shouldSubstract(value int32) bool {
	return ('N' <= value && value <= 'Z') || ('n' <= value && value <= 'z')
}

func (rot rot13Reader) Read(p []byte) (n int, err error) {
	count, readErr := rot.r.Read(p)
	
	for index, value := range p {
		if rot.shouldAdd(int32(value)) {
			p[index] = value + substitutionManipulator
		} else if rot.shouldSubstract(int32(value)) {
			p[index] = value - substitutionManipulator
		}
	}

	return count, readErr
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
