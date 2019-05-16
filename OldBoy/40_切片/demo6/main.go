package main

import "fmt"

func main() {



	m:=map[string][]int{"CVE-2017-17215":[]int{1555920345},"CVE-2017-17216":[]int{1555920345,1555921746}}
	new_m :=make(map[string]int)
	for key,value :=range m{

		if len(value)==1{
			min_value:=value
			new_m[key]=min_value[0]

		}else{
			var min int = value[0]
			for _,time := range value{
				if time<min{
					min=time
				}
            new_m[key]=min

			}
		}
		
	}
	fmt.Println(new_m)

}
