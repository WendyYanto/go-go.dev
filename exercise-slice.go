package main

import (
	"fmt"

	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	matrix := [][]uint8{} 
	for i := 0; i < dy; i++ {
		row := []uint8{}
		for j := 0; j < dx; j++ {
			row = append(row, uint8((i+j/2)))
		}
		matrix = append(matrix, row)
	}
	fmt.Println(matrix)
	return matrix
}

func main() {
	pic.Show(Pic)
}
