package main

import "sort"
import "fmt"

func main() {
	var arr = []string{"hello", "hi", "world", "hi", "china", "hello", "hi"}
	fmt.Println(RemoveRepeatedElement(arr))
}


func RemoveRepeatedElement(arr []string) (newArr []string) {
	newArr = make([]string, 0)
	sort.Strings(arr)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return
}