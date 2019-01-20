package main

import (
	"fmt"
	"gitbook/four"
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
}



