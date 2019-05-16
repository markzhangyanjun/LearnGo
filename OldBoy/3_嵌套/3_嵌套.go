package __嵌套

import "fmt"

type Com struct{
	Name string
	age  int
	class  Class
}

type Class struct{
	sex string
	cls string
	ok  []List
}

type List struct{
	OK  int
	KO  int
}

func main() {

	c := Class{"a","b",[]List{List{1,2}}}
	fmt.Println(c)
	 s :=fmt.Sprintf("c=%v",c)
	 fmt.Println(s)


	 fmt.Println(c.sex)
	 fmt.Println(c.cls)
	 fmt.Println(c.ok)
	 fmt.Println(c.ok[0])



}
