package six

import "fmt"

// Go 里面有三种类型的函数
//  普通的带有名字的函数
//  匿名函数或者lambda函数
//  方法
var (
	num int = 10
	numx2 int
	numx3 int
)

func FuncTest()  {
	g()
	a := MultiPly3Nums(1,3,4)
	fmt.Println(a)
	numx2, numx3 = getX2AndX3(num)
	PrintValue()
	numx2, numx3 = getX2AndX3_2(num)
	PrintValue()

	n := 0
	reply := &n
	Multiply(10, 5, reply)
	fmt.Println(*reply)
}

func g()  {
	print("this func name g")
}

func MultiPly3Nums(a, b, c int) int {
	return a * b * c;
}

// 如果一个函数需要返回4到5个值，我们可以船体一个切片给函数，如果返回值具有相同类型
// 或者是传递一个结构体，如果返回值具有不同的类型
// 因为传递一个指针允许直接修改变量的值，消耗也更少
/*type A struct {
	p1 int
	p2 string
}

func DoSomething(a *A)  {
	b = a
}

func DoSomething2(a A)  {
	b = &a
}*/
// 以上两个函数的区别？


func PrintValue()  {
	fmt.Printf("num= %d, 2x num = %d, 3x num = %d", num, numx2, numx3)
	fmt.Println()
}

// return 或 return var 都是可以的
// 不过 return var = expression （表达式） 会引发一个编译错误
func getX2AndX3(input int) (int, int)  {
	return 2 * input, 3 * input;
}

func getX2AndX3_2(input int) (x2 int, x3 int) {
	x2 = 2 * input
	x3 = 3 * input
	return
}

// 改变外部变量
// this function changes reply:
func Multiply(a, b int, reply *int) {
	*reply = a * b
}