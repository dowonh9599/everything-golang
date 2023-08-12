package functions

import (
	"fmt"
	"strings"
)

/*
The Foo function takes two integers a and b as arguments and returns their product.
* @param a - The first integer to be multiplied.
* @param b - The second integer to be multiplied.
* @return The product of the two integers.
*/
// Note: Go Function must specify all types to arguments and also the return type if function returns something
func Foo(a, b int) int {
	return a * b
}

/*
The LenAndUpper function takes a string name as an argument and returns its length and an uppercase version of the string.
@param name The string whose length and uppercase version are to be returned.
@return The length of the string and its uppercase version.
*/

// Note: Go Function can return multiple variables
func LenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name) // strings: Go packages for strings
}

/**
The lenAndLower function takes a string name as an argument and returns its length and a lowercase version of the string.
@param name The string whose length and lowercase version are to be returned.
@return length The length of the string.
@return lowercase The lowercase version of the string.
*/

// Note: naked return syntax of Go
func LenAndLower(name string) (length int, lowercase string) {
	defer fmt.Println("done! ") // defer executes function after the return
	length = len(name)
	lowercase = strings.ToUpper(name)
	// naked return, no need to return variables
	return
}
