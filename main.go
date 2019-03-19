package main

import (
	"fmt"
	"gitbook/four"
	"gitbook/five"
	"gitbook/six"
	"gitbook/seven"
	"strconv"
)

func main() {
	fmt.Println("Hello World")
	var ex = new(four.Example)
	ex.A = 20
	ex.B = "Fan"
	fmt.Println(ex)
	four.TestFunc()
	four.StrToAthor()
	four.TimeAndDate()
	four.PointTest()
	five.IfTest()
	five.ErrorsTest()
	six.FuncTest()
	six.FuncTest2()
	six.FuncTest3()
	six.FuncTest4()
	six.HandleFib()
	// 闭包函数
	func() {
		sum := 0
		for i := 1; i <= 1e6; i ++ {
			sum += i
		}
	}()

	// 值转换
	var in = 65
	intostr := string(in)
	fmt.Println(intostr) // A
	intostr2 := strconv.Itoa(in)
	fmt.Println(intostr2)

	// 数组
	seven.Array1()
	// 切片
	seven.Test()
	seven.Test2()
	seven.Test3()
	seven.Test4()
	seven.Test5()
	seven.Test6()
	seven.Test7()

	// map
	seven.MapTest1()
	seven.MapTest2()
	seven.MapTest3()
	seven.MapTest4()


}
