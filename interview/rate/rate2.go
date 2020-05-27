package main

import (
	"time"
	"sync"
	"fmt"
)

type LimitRate2 struct {
	rate       int
	interval   time.Duration
	lastAction time.Time
	lock       sync.Mutex
}



func (l *LimitRate2) Limit() bool {
	result := false
	for {
		l.lock.Lock()
		//判断最后一次执行的时间与当前的时间间隔是否大于限速速率
		if time.Now().Sub(l.lastAction) > l.interval {
			l.lastAction = time.Now()
			result = true
		}
		l.lock.Unlock()
		if result {
			return result
		}
		time.Sleep(l.interval)
	}
}

//SetRate 设置Rate
func (l *LimitRate2) SetRate(r int) {
	l.rate = r
	l.interval = time.Microsecond * time.Duration(1000*1000/l.rate)
}

//GetRate 获取Rate
func (l *LimitRate2) GetRate() int {
	return l.rate
}

func main() {
	var wg sync.WaitGroup
	var lr LimitRate2
	lr.SetRate(2)

	b:=time.Now()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			if lr.Limit() {
				fmt.Println("Got it!")
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(time.Since(b))
}