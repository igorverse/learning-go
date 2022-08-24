// By default cahnnels are unbuffered, meaning that they will
// only acpet sends (chan <-) if there is a corresponding
// receive (<- chan) ready to receive the sent value.
// Buffered channels accept a limited number of values
// without a corresponding receiver for those values.

package main

import "fmt"

func main() {
	// a channel of strings buffering up to 2 values
	messages := make(chan string, 2)

	// the channel is buffered, so we can send these values
	// into the channel without a correponding concurrent receive
	messages <- "buffered"
	messages <- "channel"

	// Later we can receive these two values as usual.
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}