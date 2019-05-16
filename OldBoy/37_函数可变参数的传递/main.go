package main
import "fmt"

/*
1. 0个或多个参数
	func add(arg…int) int {
	}
2. 1个或多个参数
	func add(a int, arg…int) int {
	}

3. 2个或多个参数
	func add(a int, b int, arg…int) int {
	}

注意：其中arg是一个slice，我们可以通过arg[index]依次访问所有参数
通过len(arg)来判断传递参数的个数
*/


//写一个函数add，支持1个或多个int相加，并返回相加结果
func add(a int,arg ...int)(sum int){
	if len(arg) == 0{
		sum = a
		return sum
	}else{
		for _, v :=range arg{
			a += v
		}
		sum = a
		return sum
	}

}
//写一个函数concat，支持1个或多个string相拼接，并返回结果

func concat(str string,arg ...string)(result string){
	if len(arg) ==0{
		result = str
		return str
	}else{
		for _ ,value :=range arg{
			str +=value
		}
		result = str
		return result

	}
}




func main() {

	sum := add(1)
	fmt.Println(sum)
	sum1 := add(1,2,3,4,5,6)
	fmt.Println(sum1)

	result := concat("a")
	fmt.Println(result)
	result1 := concat("a","b","c","d")
	fmt.Println(result1)

}
