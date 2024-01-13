package main

import (
	"fmt"
	"strconv"
)

func main() {
	f3 := fizzBuzz(100)
	fmt.Println(f3)
}
func fizzBuzz(n int) []string {
	var s []string
	for i := 1; i <= n; i++ {
		s = append(s, find(i))

	}
	return s
}
func find(i int) string {
	switch {
	case i%3 == 0 && i%5 == 0:
		return "FizzBuzz"

	case i%3 == 0:
		return "Fizz"

	case i%5 == 0:
		return "Buzz"
	default:
		return strconv.Itoa(i)
	}

}
