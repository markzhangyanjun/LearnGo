package main

import "time"
import "fmt"

func main() {
	second := time.Now().Unix  //1970年到现在的时间戳
	fmt.Println(second)
}
