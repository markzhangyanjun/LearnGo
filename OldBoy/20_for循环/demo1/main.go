package main

import "fmt"

func main() {
	str := "hello world"

	for index ,value :=range str{
		fmt.Printf("index[%d] value[%c] len[%d]\n",index,value,len([]byte(string(value))))

	}


}
