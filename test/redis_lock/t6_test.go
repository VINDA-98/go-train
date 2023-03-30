package redis_lock

import (
	"fmt"
	"testing"
	"time"
)

var ch = make(chan int)

func TestT6(t *testing.T) {
	go closeTest(ch)
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
	time.Sleep(time.Second * 10)
}

func closeTest(ch chan int) bool {
	for {
		select {
		case a, ok := <-ch:
			if ok {
				fmt.Println("ok：", a)
			} else {
				fmt.Println("closed")
				return false
			}
		}
	}
	return true
}
