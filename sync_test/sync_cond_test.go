//关于测试Mutex,WaitGroup,Channel
//函数调用时传值方式
//
package sync_test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var (
	locker = new(sync.Mutex)
	cond   = sync.NewCond(locker)
)
var waitGroup sync.WaitGroup

//最小业务单元
func test(x int) {

	defer func() {
		cond.L.Unlock()
		waitGroup.Done()
	}()
	cond.L.Lock()
	//waiting for notify
	cond.Wait()
	fmt.Println(x)
	time.Sleep(time.Second * 1)

}

func TestSyncCond(t *testing.T) {
	// defer func() {

	// 	if err := recover(); err != nil {
	// 		fmt.Println(err)
	// 	}
	// }()

	for i := 0; i < 5; i++ {
		waitGroup.Add(1)
		go test(i)
	}
	fmt.Println("Start All")
	time.Sleep(time.Second * 1)
	//下发一个notify
	cond.Signal()
	time.Sleep(time.Second * 1)
	//下发一个notify
	cond.Signal()
	time.Sleep(time.Second * 1)
	//下发广播
	fmt.Println("Start Boradcast")
	cond.Broadcast()
	waitGroup.Wait() //waiting for all groutines
}

func TestChannel(t *testing.T) {
	var buffer = 10
	var ch = make(chan string, buffer) //set buffer 2
	var wg = sync.WaitGroup{}
	fmt.Println("Start All")
	for i := 0; i < buffer; i++ {
		wg.Add(1)
		go func(x int, chl chan string, wa *sync.WaitGroup) {
			chl <- fmt.Sprintf("%p,%T,%v", chl, chl, x)
			wa.Done()
		}(i, ch, &wg)

	}
	wg.Wait()
	close(ch)

	//循环获取chan输出值
	for {
		val, isOpen := <-ch
		if !isOpen {
			fmt.Printf("channel[%T] is close,address:%p\n", ch, ch)
			break
		}
		fmt.Println(val)
	}
	fmt.Println("End All")
}
