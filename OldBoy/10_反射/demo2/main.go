package  main
import (
	"fmt"
	"reflect"
)

func Test(a interface{}){
	t :=reflect.TypeOf(a)
	fmt.Println(t)

	v := reflect.ValueOf(a)
	fmt.Println(v)


}

func main() {

	var a int =10

	Test(a)

}
