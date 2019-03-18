package six

// func Add2() (func(b int) int)
// func Adder(a int) (func(b int) int)

func Add2() func(b int) int {
	return func(b int) int {
		return b + 2
	}
}

func Adder(a int) func(b int) int {
	return func(b int) int {
		return a + b
	}
}
