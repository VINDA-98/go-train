package redis_lock

import (
	"fmt"
	"testing"
)

func b() {
	for i := 0; i < 4; i++ {
		defer fmt.Println(i)
	}
}

func TestT1(t *testing.T) {
	b()
}
