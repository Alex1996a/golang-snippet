package channel

import (
	"fmt"
	"testing"
)

func produceSquares(c chan int) {
	for i := 0; i <= 9; i++ {
		c <- i * i
	}

	close(c) // close channel
}

func TestReceiveChan(t *testing.T) {
	fmt.Println("main() started")
	c := make(chan int)

	go produceSquares(c) // start goroutine

	// periodic block/unblock of main goroutine until chanel closes
	for val := range c {
		fmt.Println(val)
	}

	fmt.Println("main() stopped")
}
