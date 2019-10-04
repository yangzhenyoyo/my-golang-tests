package sync_test
import(
	"testing"
	"time"
	"fmt"
	"sync"
)

func isCanceled(ch chan struct{})bool{
	select{
	 case <-ch:
		 return true
	 default:
		 return false	 
	}
}

func cancel_1(ch chan struct{}){
	  ch<- struct{}{}
}

func cancel_2(ch chan struct{}){
	close(ch)
}

func TestMutil(t *testing.T){
		ch :=make(chan struct{})
	  var mwg = &sync.WaitGroup{}
	  for i:=1;i<5;i++{
			mwg.Add(1)
			 go func(i int, ch2 chan struct{},wg *sync.WaitGroup){
           for j:=0;j<5;j++{
						if(isCanceled(ch2)){
							fmt.Println(i,"is canceled")
							break;
						}
						fmt.Println(i,"is working,try time:",j)
						time.Sleep(time.Millisecond*5)


					 }
					 
					 wg.Done()

			 }(i,ch,mwg)
		} 
	 
		// cancel_1(ch)
		time.Sleep(time.Millisecond*10)
	cancel_1(ch)  
		mwg.Wait()
	
}