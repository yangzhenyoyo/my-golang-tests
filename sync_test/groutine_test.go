package sync_test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func show(n int) {
	fmt.Println("the value is:", n)
}

func TestGroutine(t *testing.T) {
	for i := 0; i <= 10; i++ {
		go show(i)
	}

	time.Sleep(time.Microsecond * 50)

}

func TestWaitGroup(t *testing.T) {
	var mt sync.Mutex
	var wg sync.WaitGroup
	count := 0
	for i := 0; i < 5000; i++ {

		go func() {
			wg.Add(1)
			defer func() {
				mt.Unlock()
			}()
			mt.Lock()
			count++
			wg.Done()
		}()

	}
	wg.Wait()
	fmt.Println("count is ", count)

}
