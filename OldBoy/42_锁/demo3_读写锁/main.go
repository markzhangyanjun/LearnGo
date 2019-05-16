package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

//var rwLock sync.RWMutex
var lock sync.Mutex
func testMap(){

	a :=make(map[int]int,5)
	var count int32
	a[1]=10
	a[2]=10
	a[3]=10
	a[4]=10
	a[5]=10

	for i:=0;i<2;i++{
		go func(b map[int]int){
			//rwLock.Lock()
			lock.Lock()
			b[1]=rand.Intn(100)
			//rwLock.Unlock()
			lock.Unlock()


		}(a)
	}

	for i:=0;i<100;i++{
		go func(b map[int]int){
			for{
				lock.Lock()
				//rwLock.RLock()
				fmt.Println(a)
				//rwLock.RUnlock()
				lock.Unlock()
				atomic.AddInt32(&count,1)

			}

		}(a)
	}
    time.Sleep(time.Second*3)
	fmt.Println(atomic.LoadInt32(&count))


}



func main() {
	testMap()
}
