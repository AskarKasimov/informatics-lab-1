package main

import (
	"fmt"
	"slices"
	"strconv"
)

func getFibs(target int) []int {
	var fibs = []int{1, 2} // declaring slice with fibonacci numbers

	i := 2
	for {
		newFib := fibs[i-2] + fibs[i-1]
		if newFib > target { // looking for the biggest fib lower than given
			break
		}
		fibs = append(fibs, newFib)
		i += 1
	}

	return fibs
}

func greedy(a int, fibs []int) string {
	result := ""
	for _, value := range fibs {
		if a-value >= 0 { // looking for the largest fib to subtract
			a -= value

			result += "1" // forming the result string based on fib index
		} else {
			result += "0"
		}
	}
	return result
}

func FromDecToFib(input int) string {
	fibs := getFibs(input)
	slices.Reverse(fibs) // reorganazing slice to place the biggest fibs at the beggining
	return greedy(input, fibs)
}

func main() {
	var input string
	fmt.Scanln(&input)
	number, _ := strconv.Atoi(input)
	fmt.Println(FromDecToFib(number))
}
