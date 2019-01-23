package six

import (
	"fmt"
	"log"
	"io"
)

func FuncTest4()  {
	b()
	func1("GODefer")
}
func tr(s string) string {
	fmt.Println("entering:", s)
	return s
}

func un(s string) {
	fmt.Println("leaving:", s)
}

func a2() {
	defer un(tr("a"))
	fmt.Println("in a")
}

func b() {
	defer un(tr("b"))
	fmt.Println("in b")
	a2()
}

/*
entering: b
in b
entering: a
in a
leaving: a
leaving: b
*/

func func1(s string) (n int, err error) {
	defer func() {
		log.Printf("func1(%q) = %d, %v", s, n, err)
	}()
	return 7, io.EOF
}
// 2019/01/24 00:05:38 func1("GODefer") = 7, EOF