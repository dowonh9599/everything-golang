package conditionalIterative

import "fmt"

/*
The factorial function takes an integer number as an argument and returns its factorial value.
@param number The integer whose factorial value is to be calculated.
@return The factorial value of the given integer.
*/
// Note: for-loop syntax in Go
func Factorial(number int) (res int) {
	if number < 1 {
		error := fmt.Errorf("number cannot be less than 1")
		fmt.Println(error)
		return 0
	}
	res = number
	for i := number - 1; i > 0; i-- {
		res *= i
	}
	return
}

/*
The sum function takes one or more integers as arguments and returns their sum.
@param numbers The integers whose sum is to be calculated.
@return The sum of the given integers.
*/
// Note: shorthand form of for loop
func Sum(numbers ...int) (total int) {
	for index, number := range numbers {
		fmt.Println("current total", total)
		fmt.Println("adding", numbers[index])
		total += number
		fmt.Println("result", total)
		fmt.Println()
	}
	return
}
