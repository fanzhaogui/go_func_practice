package six

import "fmt"

func HandleFib()  {
	result := 0
	for i:= 0; i <= 10; i++ {
		result = Fib1(i)
		fmt.Println(result)
	}
	FuncClourse()

	p2 := Add2()
	fmt.Printf("Call Add2 for 3 gives:%v\n", p2(3))
	TwoAdder := Adder(2)
	fmt.Printf("The result is: %v\n", TwoAdder(3))
}

func Fib1(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = Fib1(n - 1) + Fib1(n - 2)
	}
	return
}