package goroutine_test

import (
	"fmt"
	"testing"
)

func chan_1() {
	var ch = make(chan int, 10)
	ch <- 1
	ch <- 2
	fmt.Println(len(ch), cap(ch))
}

func TestT1(t *testing.T) {
	chan_1()
}
