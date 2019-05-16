package main
import "fmt"
func main() {

	s1 :=[]int{1,2,3,4}
	s2:= make([]int,10)
	copy(s2,s1)
	fmt.Println(s2)

	s3:=make([]int,0)
	s3 = append(s3,1)
	s3 = append(s3,2,3,4,5)
	s4:= []int{7,8,9}
	s3= append(s3,s4...)
	fmt.Println(s3)

}
