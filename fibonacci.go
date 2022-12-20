package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	firstNumber := 0
	secondNumber := 1

	invokeCount := 0

	return func() int {
		defer func() {
			invokeCount++
		}()
		if invokeCount == 0 {
			return firstNumber
		} else if invokeCount == 1 {
			return secondNumber
		}

		result := firstNumber + secondNumber

		firstNumber = secondNumber
		secondNumber = result

		return result
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
