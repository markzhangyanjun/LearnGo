package demo1
import "fmt"
/*
1. 切片：切片是数组的一个引用，因此切片是引用类型
2. 切片的长度可以改变，因此，切片是一个可变的数组
3. 切片遍历方式和数组一样，可以用len()求长度
4. cap可以求出slice最大的容量，0 <= len(slice) <= cap(array)，其中array是slice引用的数组
5. 切片的定义：var 变量名 []类型，比如 var str []string  var arr []int
*/

func main() {
	var s [4]int=[...]int{1,2,3,4}
	fmt.Println(s)
}
