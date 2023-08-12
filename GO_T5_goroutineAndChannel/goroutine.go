package goroutineAndChannel

import (
	"fmt"
	"time"
)

func TestGoroutine(withGoKeyword bool) {
	// Note: counter alpha and beta counts concurrently by adding keyword "go" to task
	if withGoKeyword {
		// this will have counter alpha and beta running together at the background
		counters := [2]string{"alpha", "beta"}
		for _, counter := range counters {
			go printCount(counter)
		}
		// Important:
		// - should have at least one task for function to run, Go functions does not wait for goroutine to finish the task, goroutine only lives while function has tasks.
		time.Sleep(time.Second * 10)
	} else {
		counters := [2]string{"alpha", "beta"}
		// this will first run counter alpha from 0-9, then run the counter beta
		for _, counter := range counters {
			printCount(counter)
		}
	}
}

func TestChannel() {
	// create a channel
	// specify the type of the data channel will return
	c := make(chan string)
	fmt.Println("waiting for the message to arrive after 5 seconds...")
	messages := [2]string{"hello world!", "bye world!"}

	// add two goroutines at the background
	for _, message := range messages {
		go sendMessageThroughChannel(message, c)
	}
	// get messages from channel

	// notice that this function has no time sleep, but the function waits 5 seconds implemented at the "sendMessageThroughChannel"
	// two lines of codes below are the blocking operation: meaning will make the function to wait for response to arrive
	// but having three lines of waiting for response will create an error (deadlock): all groutines are asleep
	// - since there's only two goroutines registered at the back

	// Tip: no need to hardcode line by line, can include in loop
	msg1 := <-c
	msg2 := <-c

	fmt.Println("message arrived: ", msg1)
	fmt.Println("message arrived: ", msg2)
}

func printCount(label string) {
	for i := 0; i < 10; i++ {
		fmt.Println(label, "...", i)
		time.Sleep(time.Second)
	}
}

func sendMessageThroughChannel(msg string, c chan string) {
	time.Sleep(time.Second * 5)
	c <- msg
}
