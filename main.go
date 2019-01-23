package main

import (
	"fmt"
	"gitbook/four"
	"gitbook/five"
	"gitbook/six"
)

func main()  {
	fmt.Println("Hello World")
	var ex = new (four.Example)
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
}



