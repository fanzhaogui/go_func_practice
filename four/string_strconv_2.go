package four

import (
	"fmt"
	"reflect"
	"strconv"
)

// 字符串与其他类型的转换
func StrToAthor()  {
	var orgS = 2;
	// Itoa 返回数字i所表示的字符串类型的十进制数
	fmt.Println(reflect.TypeOf(orgS))
	fmt.Println(reflect.TypeOf(strconv.Itoa(orgS)))


	// strconv.Atoi(s string) (i int, err error) 将字符串转换为 int 型。
	var orgI = "666"
	fmt.Println(strconv.Atoi(orgI))

	// strconv.ParseFloat(s string, bitSize int) (f float64, err error) 将字符串转换为 float64 型。
	var orgF = "12.56"
	var orgTfloate float64
	orgTfloate , _= strconv.ParseFloat(orgF, 2)
	fmt.Println(orgTfloate)
	// strconv.FormatFloat(f float64, fmt byte, prec int, bitSize int) 将64位浮点型数字转换为字符串


}
