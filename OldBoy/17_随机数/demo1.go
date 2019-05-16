package main

import (
	"fmt"
	"math/rand"
)

func main() {
	a :=rand.Int()
	fmt.Println(a)

	b := rand.Intn(100 )
	fmt.Println(b)
}
