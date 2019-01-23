package six

import (
	"fmt"
)

func FuncTest2()  {
	myFunc(1,2,3,4,5)

	min1 := min(8,23,4,5,6,889,95)
	fmt.Println(min1)
	//slice := []int{4,3,5,6,7}
	//x := min(slice) // 如法传值？

	function1()
	a()
	f()
}
// 传递边长参数
// 如果函数的最后一个参数是采用的...type的形式，name这个函数就可以处理一个变长的参数，这个长度可以为0
func myFunc(a, b int, arg ...int)  {
	fmt.Println(a,b,arg)
}

func min(s ...int) int {
	// fmt.Println(reflect.Type(s))
	if len(s) == 0 {
		return 0
	}
	min := s[0]
	for _,v := range s {
		if v < min {
			min = v
		}
	}
	return  min
}

// 关键字defer 允许我们推迟到函数返回之前（或任意位置执行return语句之后）一刻才执行某个语句或函数
// 为什么要在返回只有才执行这些语句？
// 因为return语句同样可以包含一些操作，而不是单纯地返回某个值
func function1()  {
	fmt.Println("In function1 at the top")
	//defer function2()
	function2() // 对比去掉defer的输出结果
	fmt.Println("In function2 at the bottom")
}

func function2()  {
	fmt.Println("Function: Deferred until the end of the calling function")
}

func a()  {
	i := 0
	defer fmt.Println(i)
	i ++
	return
}

// 当有多个defer行为被注册司，他们会以逆序执行(类似栈，即后进先出)
func f()  {
	for i:=0; i<5; i++ {
		defer fmt.Println(i) // 4 3 2 1 0
	}
}