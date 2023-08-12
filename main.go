package main

import (
	"fmt"
	"github.com/dowonh9599/go-basics/GO_E2_dictionary"
	"github.com/dowonh9599/go-basics/GO_T5_goroutineAndChannel"
)

func main() {
	// exercise 1 banking

	// exercise 2 Dictionary
	dictionary := dict.CreateDictionary()
	addAppleErr := dictionary.Add("apple", "the round fruit of a tree of the rose family, which typically has thin green or red skin and crisp flesh.")
	if addAppleErr != nil {
		fmt.Println(addAppleErr)
	}
	addBananaErr := dictionary.Add("banana", "a long curved vegetable which grows in clusters and has soft pulpy flesh and yellow skin when ripe.")
	if addBananaErr != nil {
		fmt.Println(addBananaErr)
	}
	updateBananaErr := dictionary.Update("banana", "a long curved fruit which grows in clusters and has soft pulpy flesh and yellow skin when ripe.")
	if updateBananaErr != nil {
		fmt.Println(updateBananaErr)
	}
	errDeleteBanana := dictionary.Delete("banana")
	if errDeleteBanana != nil {
		fmt.Println(errDeleteBanana)
	}
	dictionary.PrintDictionary(dictionary)

	// goroutineAndChannel exercise
	goroutineAndChannel.TestGoroutine(false)
	goroutineAndChannel.TestChannel()

	// Compare the processing speed of URLChecker handling 10 URL checks using / not using goroutine
	goroutineAndChannel.TestURLChecker()
	goroutineAndChannel.TestURLCheckerWithGoroutine()

}
