package six

import "fmt"

func FuncTest3()  {
	f2()
}

func trace(s string) {
	fmt.Println("entering", s)
}

func untrace(s string)  {
	fmt.Println("leaving", s)
}

func f1()  {
	trace("f1")
	defer untrace("f1")
	fmt.Println("in f1")
}

func f2()  {
	trace("f2")
	defer untrace("f2")
	fmt.Println("in f2")
	f1()
}
/*
entering f2
in f2
entering f1
in f1
leaving f1
leaving f2
*/