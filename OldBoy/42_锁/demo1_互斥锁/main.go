package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//互斥锁
var lock sync.Mutex


func testMap(){

	a :=make(map[int]int,5)
	a[1]=10
	a[2]=10
	a[3]=10
	a[4]=10
	a[5]=10

	for i:=0;i<2;i++{
		go func(b map[int]int){
			lock.Lock()
			b[1]=rand.Intn(100)
			lock.Unlock()

		}(a)
	}
	time.Sleep(time.Second)
    lock.Lock()
	fmt.Println(a)
	lock.Unlock()


}



func main() {
	testMap()
}

