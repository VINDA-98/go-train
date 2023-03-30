package redis_lock

import (
	"sync"
	"testing"
	"time"

	"github.com/go-redis/redis"
)

// 全局变量
var counter int
var counter1 int
var counter2 int

func waitFailed() {
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter++
		}()
	}

	wg.Wait()
	println(counter)
}

func TestT4(t *testing.T) {
	//waitFailed()
	//waitProcess()
	//var lock = NewLock()
	//var wg sync.WaitGroup
	//for i := 0; i < 10; i++ {
	//	wg.Add(1)
	//	go func() {
	//		defer wg.Done()
	//		if !lock.Lock() {
	//			println("lock failed")
	//			return
	//		}
	//		counter2++
	//		println("counter2:", counter2)
	//		lock.Unlock()
	//	}()
	//}
	//wg.Wait()
	//waitRedis()
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			setnx()
		}()
	}
	wg.Wait()

}

// 在进程内加锁
func waitProcess() {
	var wg sync.WaitGroup
	var lock sync.Mutex
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			lock.Lock()
			counter1++
			lock.Unlock()
		}()
	}
	wg.Wait()
	println(counter1)
}

type Lock struct {
	c chan struct{}
}

func (l Lock) Lock() bool {
	result := false
	select {
	case <-l.c:
		result = true
	default:

	}

	return result
}

// Unlock the try lock
func (l Lock) Unlock() {
	l.c <- struct{}{}
}

func NewLock() Lock {
	var l Lock
	l.c = make(chan struct{}, 1)
	l.c <- struct{}{}
	return l
}

func setnx() {
	client := redis.NewClient(&redis.Options{})

	var lockKey = "counter_lock"
	var counterKey = "counter"

	// lock
	resp := client.SetNX(lockKey, 1, time.Second*5)
	lockStatus, err := resp.Result()
	println("lockStatus:", lockStatus)
	if err != nil || !lockStatus {
		println("lock failed")
		return
	}

	// counter++
	getResp := client.Get(counterKey)
	cntValue, err := getResp.Int64()
	if err == nil || err == redis.Nil {
		cntValue++
		resp := client.Set(counterKey, cntValue, 0)
		_, err := resp.Result()
		if err != nil {
			println(err)
		}
	}
	println("current counter is ", cntValue)

	// unlock
	delResp := client.Del(lockKey)
	unlockStatus, err := delResp.Result()
	if err == nil && unlockStatus > 0 {
		println("unlock success")
	} else {
		println("unlock failed", err)
	}
}

func waitRedis() {

}
