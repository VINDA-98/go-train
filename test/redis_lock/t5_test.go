package redis_lock

import (
	"fmt"
	"testing"
)

func TestT5(t *testing.T) {
	def_call() //后进先出

}

func def_call() {
	defer func() {
		fmt.Println("打印前")
	}()
	defer func() {
		fmt.Println("打印中")
	}()
	defer func() {
		fmt.Println("打印后")
	}()
	panic("触发异常")
}
