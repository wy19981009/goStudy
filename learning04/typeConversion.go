package main

import (
	"fmt"
	"strconv"
)

func typeConversion() {
	// 字符串与其他类型的转换
	// str 转 int
	newStr1 := "1"
	intValue, _ := strconv.Atoi(newStr1)
	fmt.Printf("%T,%d\n", intValue, intValue) // int,1

	// int 转 str
	intValue2 := 1
	strValue := strconv.Itoa(intValue2)
	fmt.Printf("%T, %s\n", strValue, strValue)

	// str 转  float
	string3 := "3.1415926"
	f, _ := strconv.ParseFloat(string3, 32)
	fmt.Printf("%T, %f\n", f, f) // float64, 3.141593
	//float 转 string
	floatValue := 3.1415926
	//4个参数，1：要转换的浮点数 2. 格式标记（b、e、E、f、g、G）
	//3. 精度 4. 指定浮点类型（32:float32、64:float64）
	// 格式标记：
	// ‘b’ (-ddddp±ddd，二进制指数)
	// ‘e’ (-d.dddde±dd，十进制指数)
	// ‘E’ (-d.ddddE±dd，十进制指数)
	// ‘f’ (-ddd.dddd，没有指数)
	// ‘g’ (‘e’:大指数，‘f’:其它情况)
	// ‘G’ (‘E’:大指数，‘f’:其它情况)
	//
	// 如果格式标记为 ‘e’，‘E’和’f’，则 prec 表示小数点后的数字位数
	// 如果格式标记为 ‘g’，‘G’，则 prec 表示总的数字位数（整数部分+小数部分）
	formatFloat := strconv.FormatFloat(floatValue, 'f', 2, 64)
	fmt.Printf("%T,%s", formatFloat, formatFloat)
}
