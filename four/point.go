package four

import "fmt"

//  指针

func PointTest()  {
	var i1 = 5
	fmt.Printf("An integer: %d, it's location in memory: %p\n", i1, &i1)

	var intP *int
	intP = &i1
	fmt.Println(intP)
}
