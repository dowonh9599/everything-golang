package functions

import (
	"fmt"
	"strings"
)

// in Go, capitalize the first case of function name to export it (public function)
func SayHello() {
	fmt.Printf("Hello!")
}

func SayBye() {
	fmt.Printf("Bye!")
}

func askHowAreYou() {
	fmt.Printf("How are you?")
}

func RepeatStrings(words ...string) string {
	return strings.Join(words, " ")
}
