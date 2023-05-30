package main

import (
	"context"
	"log"
	"time"
)

// @Title  context_test
// @Description  MyGO
// @Author  WeiDa  2023/5/30 10:18
// @Update  WeiDa  2023/5/30 10:18

func main() {
	ctx, cancelFunc := context.WithCancel(context.Background())
	go Swimming(ctx)
	time.Sleep(3 * time.Second)
	cancelFunc()
	time.Sleep(1 * time.Second)
}

func Swimming(ctx context.Context) {
	for {
		select {
		case <-ctx.Done(): //接收到停止游泳的信号
			log.Println("别游泳了,歇会吧...")
			return
		default:
			log.Println("正在游泳...")
		}
	}
}
