package main

import "fmt"

func main() {

	s := []string{"a","b","c"}
	m2 := make(map[string]string)
	for index,v :=range s{
		fmt.Println(index,v)
		m2[v]=v

	}
	fmt.Println(m2)

	result,err:=m2["f"]
	fmt.Println("result:=",result)
	fmt.Println(err)

	}

