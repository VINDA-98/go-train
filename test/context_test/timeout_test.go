package main

import (
	"context"
	"log"
	"testing"
	"time"
)

// @Title  main
// @Description  MyGO
// @Author  WeiDa  2023/5/30 11:01
// @Update  WeiDa  2023/5/30 11:01

func TestT1(t *testing.T) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 3*time.Second)

	go func() {
		time.Sleep(5 * time.Second) //模拟数据库业务，此时已经超时
		cancelFunc()
	}()

	defer cancelFunc()
	select {
	case <-ctx.Done(): //接收到完成信号，因为已经过了五秒，操作超时
		log.Println("操作已超时")
	case <-time.After(10 * time.Second):
		log.Println("操作已完成")
	}

}
